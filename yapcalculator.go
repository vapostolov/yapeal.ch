package yapstones

import "fmt"

// YapCalculator contains calculators used in amount and money
type YapCalculator struct{}

const mostNegative = -(mostPositive + 1)
const mostPositive = 1<<63 - 1

// Add to amounts
func (c *YapCalculator) Add(a, b *YapAmount) (ya *YapAmount, err error) {
	if a.Factor == b.Factor {
		ya = &YapAmount{Value: a.Value + b.Value, Factor: a.Factor}
	} else {
		a.Normalize()
		b.Normalize()
		ya = &YapAmount{Value: a.Value + b.Value, Factor: a.Factor}
	}
	return
}

// Subtract two amounts
func (c *YapCalculator) Subtract(a, b *YapAmount) (ya *YapAmount, err error) {
	if a.Factor == b.Factor {
		ya = &YapAmount{Value: a.Value - b.Value, Factor: a.Factor}
	} else {
		a.Normalize()
		b.Normalize()
		ya = &YapAmount{Value: a.Value - b.Value, Factor: a.Factor}
	}
	return
}

// IsEqual determines equality
func (c *YapCalculator) IsEqual(a, b *YapAmount) (eq bool) {

	if a.Factor == b.Factor {
		eq = a.Value == b.Value
	} else {
		a.Normalize()
		b.Normalize()
		eq = a.Value == b.Value
	}

	return
}

// Multiply two amounts
func (c *YapCalculator) Multiply(a, b *YapAmount) (ya *YapAmount, err error) {
	if a.Value == 0 || b.Value == 0 {
		ya = &YapAmount{Value: 0, Factor: 0}
		return
	}

	result := a.Value * b.Value
	ya = &YapAmount{Value: result, Factor: a.Factor + b.Factor}

	if a.Value == 1 || b.Value == 1 {
		return
	}

	if a.Value == mostNegative || b.Value == mostNegative {
		err = fmt.Errorf("Overflow multiplying %v and %v", a, b)
		return
	}
	if result/b.Value != a.Value {
		err = fmt.Errorf("Overflow multiplying %v and %v", a, b)
		return
	}
	return
}

// Divide two amounts
func (c *YapCalculator) Divide(a, b *YapAmount, precision uint8) (ya *YapAmount, err error) {
	if b.Value == 0 {
		err = fmt.Errorf("Division by zero error")
		return
	}

	var factor uint8
	if a.Factor > b.Factor {
		factor = a.Factor
	} else {
		factor = b.Factor
	}

	a.NormalizeWith(factor)
	b.NormalizeWith(factor)
	remainder := a.Value % b.Value
	result := a.Value / b.Value
	if remainder == 0 {
		ya = &YapAmount{Value: result, Factor: 0}
		return
	}
	factor = 0
	for precision > 0 {
		result *= 10
		remainder *= 10
		result += remainder / b.Value
		remainder = remainder % b.Value
		precision--
		factor++
	}
	ya = &YapAmount{Value: result, Factor: factor}
	return
}
