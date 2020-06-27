package yapstones

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// DefaultMultiplier used if none is specified
const DefaultMultiplier = 6

// multipliers array used to shift representations
var multipliers = map[uint8]int64{
	0: 1,
	1: 10,
	2: 100,
	3: 1000,
	4: 10000,
	5: 100000,
	6: 1000000,
	7: 10000000,
	8: 100000000,
	9: 1000000000,
}

// zeros all zeros used for padding
const zeros = "0000000000000000"

// YapAmount realized an amount implemented with an int64 and a shift factor
// The shift factor can change and does not really matter. its main purpose
// is to allow the suage of intgers to store the value.
// Example: value 1000 factor 2 means 10.00

type YapAmount struct {
	Value  int64 `json:"value"`  // value multiplied by Factor
	Factor uint8 `json:"factor"` // shift factor

}

// Abs returns the absolute of the given amount
// Example: Abs(value -2000 factor 2) is value 2000 factor 2
func (y YapAmount) Abs() (a YapAmount) {
	a = y
	a.Value = Abs(y.Value)
	return
}

// AmountAsString returns the amount as a string
func (y YapAmount) AmountAsString() (res string) {

	var v int64
	var f int64
	var pad uint8

	v = Abs(y.Value) / multipliers[y.Factor]
	f = Abs(y.Value) % multipliers[y.Factor]

	vs := strconv.FormatInt(v, 10)
	fs := strconv.FormatInt(f, 10)

	if y.Factor == 0 {
		pad = 0
	} else {
		pad = y.Factor - uint8(len(fs))
	}

	if y.Value >= 0 {
		if (pad + uint8(f)) > 0 {
			res = vs + "." + zeros[0:pad] + fs
		} else {
			res = vs
		}
	} else {
		if (pad + uint8(f)) > 0 {
			res = "-" + vs + "." + zeros[0:pad] + fs
		} else {
			res = "-" + vs
		}
	}

	return
}

// AmountAsInt64 returns the integer amount
func (y *YapAmount) AmountAsInt64() (res int64) {
	res = y.Value / multipliers[y.Factor]
	return
}

func (y *YapAmount) AmountAsFloat64() (res float64) {
	res = float64(y.Value) / float64(multipliers[y.Factor])
	return
}

// AmountFromString converts the string to an amount
func (y *YapAmount) AmountFromString(value string) (err error) {

	var v int64
	var f int64

	parts := strings.Split(value, ".")
	if len(parts) == 0 {
		// no decimal point
		v, err = strconv.ParseInt(value, 0, 64)
		if err != nil {
			log.Warnf("amount conversion %v %v", value, err)
			return
		}

	} else {
		v, err = strconv.ParseInt(parts[0], 0, 64)
		if err != nil {
			log.Warnf("amount conversion %v %v", value, err)
			return
		}
		if len(parts) == 2 {

			l := len(parts[1])
			// fmt.Printf("len %v", l)

			v := parts[1]
			v = strings.TrimLeft(v, "0")
			if v == "" {
				v = "0"
			}

			f, err = strconv.ParseInt(v, 0, 64)
			// fmt.Printf("f %v", f)

			if err != nil {
				log.Warnf("amount conversion %v %v", value, err)
				return
			}
			f = f * multipliers[uint8(DefaultMultiplier-l)]
		}
	}
	y.Factor = DefaultMultiplier
	y.Value = v*multipliers[y.Factor] + f
	return
}

// AmountFromStringMultiplyer converts the string to an amount qith a specific multiplier
func (y *YapAmount) AmountFromStringMultiplyer(value string, multiplyer uint8) (err error) {

	if err = y.AmountFromString(value); err != nil {
		return
	}

	y.NormalizeWith(multiplyer)
	return
}

// Normalize the representation
// Applies the default multiplier
func (y *YapAmount) Normalize() {
	if y.Factor != DefaultMultiplier {
		if y.Factor > DefaultMultiplier {
			d := y.Factor - DefaultMultiplier
			y.Value = y.Value / multipliers[d]
		} else {
			d := DefaultMultiplier - y.Factor
			y.Value = y.Value * multipliers[d]
		}
		y.Factor = DefaultMultiplier
	}
}

// Normalize the representation to a specific multiplier
func (y *YapAmount) NormalizeWith(multiplier uint8) {
	if y.Factor != multiplier {
		if y.Factor > multiplier {
			d := y.Factor - multiplier
			y.Value = y.Value / multipliers[d]
		} else {
			d := multiplier - y.Factor
			y.Value = y.Value * multipliers[d]
		}
		y.Factor = multiplier
	}
}

// Abs determines the absolute value
func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// GetParts with 10 digits as strings
func (y *YapAmount) GetParts() (intPart, fractPart string) {

	v := Abs(y.Value) / multipliers[y.Factor]
	f := Abs(y.Value) % multipliers[y.Factor]

	vs := strconv.FormatInt(v, 10)
	fs := strconv.FormatInt(f, 10)

	pad := y.Factor - uint8(len(fs))

	if y.Value >= 0 {
		intPart = vs
	} else {
		intPart = "-" + vs
	}

	fmt.Printf("vs %v fs %v pad %v ", vs, fs, pad)
	fractPart = zeros[0:pad] + fs

	return
}
