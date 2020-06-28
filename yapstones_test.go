package yapstones

import (
	"strconv"
	"testing"
)

func TestConvversions(t *testing.T) {

	// allValid := true

	t.Log("------------------------------------------------------------------------------------------")
	{
		s1 := "100"
		var y YapAmount

		if err := y.AmountFromString(s1); err != nil {
			t.Errorf("conversion resulted in unexpected error %v", err)
		}
		s2 := y.AmountAsString()

		if s2 != "100.000000" {
			t.Errorf("Test conversions case  not ok  %v --> %v ---> %v", s1, y, s2)
		}
	}
	t.Log("------------------------------------------------------------------------------------------")
	{
		s1 := ""
		var y YapAmount

		if err := y.AmountFromString(s1); err == nil {
			t.Errorf("conversion resulted in no error ")
		}
	}
	t.Log("------------------------------------------------------------------------------------------")
	{
		s1 := "1fgfgf1.56565ghg"
		var y YapAmount

		if err := y.AmountFromString(s1); err == nil {
			t.Errorf("conversion resulted in no error ")
		}
	}
	t.Log("------------------------------------------------------------------------------------------")
	{
		s1 := "8223372036854.775807"
		var y YapAmount

		if err := y.AmountFromString(s1); err != nil {
			t.Errorf("conversion resulted in unexpected error %v", err)
		}
		s2 := y.AmountAsString()

		if s2 != "8223372036854.775807" {
			t.Errorf("Test conversions case  not ok  %v --> %v ---> %v", s1, y, s2)
		}
	}
	t.Log("------------------------------------------------------------------------------------------")
	{
		s1 := "4223372036854.775807"
		s2 := "1776627963145.224193"
		sr := "6000000000000.000000"

		var a YapAmount
		var b YapAmount
		var r YapAmount

		if err := a.AmountFromString(s1); err != nil {
			t.Errorf("conversion resulted in unexpected error %v", err)
		}
		if err := b.AmountFromString(s2); err != nil {
			t.Errorf("conversion resulted in unexpected error %v", err)
		}
		if err := r.AmountFromString(sr); err != nil {
			t.Errorf("conversion resulted in unexpected error %v", err)
		}

		var c YapCalculator

		rc, _ := c.Add(&a, &b)

		if !c.IsEqual(&r, rc) {
			rcs := rc.AmountAsString()
			t.Errorf("expected equality %v and %v ", sr, rcs)

		}

		tc, _ := c.Subtract(rc, &b)

		if !c.IsEqual(&r, rc) {
			tcs := tc.AmountAsString()
			t.Errorf("expected equality %v and %v ", s1, tcs)

		}

	}
	t.Log("------------------------------------------------------------------------------------------")
	{
		// s1 := "123456.123456"
		// s2 := "10"

		// var a YapAmount
		// var b YapAmount

		// if err := a.AmountFromString(s1); err != nil {
		// 	t.Errorf("conversion resulted in unexpected error %v", err)
		// }
		// if err := b.AmountFromString(s2); err != nil {
		// 	t.Errorf("conversion resulted in unexpected error %v", err)
		// }

		// var c YapCalculator

		// r, _ := c.Multiply(&a, &b)

		// sr := r.AmountAsString()

		// if sr != "1234561.234560" {
		// 	t.Errorf("Test multiplication case  not ok  %v --> %v ---> %v", s1, s2, sr)
		// }
	}
	t.Log("------------------------------------------------------------------------------------------")
	{

		var a YapAmount
		var b YapAmount
		var r YapAmount

		a.Value = 100
		a.Factor = 1

		b.Value = 100000
		b.Factor = 4

		r.Value = 20000000
		r.Factor = 6

		var c YapCalculator

		rc, _ := c.Add(&a, &b)

		if !c.IsEqual(&r, rc) {
			t.Errorf("expected equality %v and %v ", r.AmountAsString(), rc.AmountAsString())
		}

		tc, _ := c.Subtract(rc, &b)

		if !c.IsEqual(&a, tc) {
			t.Errorf("expected equality %v and %v ", b.AmountAsString(), tc.AmountAsString())
		}

	}
}

func TestCoversion(t *testing.T) {

	s1 := "4223372036854.775807"
	s2 := "1776627963145.224193"
	s3 := "6000000000000.000010"
	s4 := "37310.090000"

	var a YapAmount
	var b YapAmount
	var c YapAmount
	var d YapAmount

	if err := a.AmountFromString(s1); err != nil {
		t.Errorf("conversion resulted in unexpected error %v", err)
	}
	if err := b.AmountFromString(s2); err != nil {
		t.Errorf("conversion resulted in unexpected error %v", err)
	}
	if err := c.AmountFromString(s3); err != nil {
		t.Errorf("conversion resulted in unexpected error %v", err)
	}
	if err := d.AmountFromString(s4); err != nil {
		t.Errorf("conversion resulted in unexpected error %v", err)
	} else if d.AmountAsString() != s4 {
		t.Errorf("conversion unexpected value %v", d.AmountAsString())
	}

}

func TestMultiplier2(t *testing.T) {

	amounts := []string{
		"37310.09",
		"37310.9",
		"37310",
		"0.01",
		"0.10",
		"0.1",
		"100000",
	}

	expected := []string{
		"37310.09",
		"37310.90",
		"37310.00",
		"0.01",
		"0.10",
		"0.10",
		"100000.00",
	}

	for i, s := range amounts {
		var d YapAmount

		if err := d.AmountFromString(s); err != nil {
			t.Errorf("conversion resulted in unexpected error %v", err)
		} else {
			d.NormalizeWith(2)
			result := d.AmountAsString()
			if result != expected[i] {
				t.Errorf("conversion value %v expected %v received %v", s, expected[i], result)
			}
		}
	}

	for i, s := range amounts {
		var d YapAmount

		if err := d.AmountFromStringMultiplyer(s, 2); err != nil {
			t.Errorf("conversion resulted in unexpected error %v", err)
		} else {
			result := d.AmountAsString()
			if result != expected[i] {
				t.Errorf("conversion value %v expected %v received %v", s, expected[i], result)
			}
		}
	}
}

func TestComparison(t *testing.T) {

	var a, b YapAmount
	var c YapCalculator

	a.Value = 1000
	a.Factor = 2

	b.Value = 100
	b.Factor = 1

	if !c.IsEqual(&a, &b) {
		as := a.AmountAsString()
		bs := b.AmountAsString()
		t.Errorf("Expected equality  %v ==  %v", as, bs)
	}

	b.Value = 1000

	if c.IsEqual(&a, &b) {
		as := a.AmountAsString()
		bs := b.AmountAsString()
		t.Errorf("Expected no equality  %v ==  %v", as, bs)
	}

	b.Value = 10000
	b.Factor = 3

	if !c.IsEqual(&a, &b) {
		as := a.AmountAsString()
		bs := b.AmountAsString()
		t.Errorf("Expected equality  %v ==  %v", as, bs)
	}

}

func TestParts(t *testing.T) {

	var a YapAmount
	var b YapAmount
	var c YapAmount

	a.Value = 101
	a.Factor = 1

	b.Value = 12340001
	b.Factor = 4

	c.Value = -101
	c.Factor = 1

	ipa, fpa := a.GetParts()

	if ipa != "10" {
		t.Errorf("unexpected result %v ", ipa)
	}
	if fpa != "1" {
		t.Errorf("unexpected result %v ", fpa)
	}

	ipb, fpb := b.GetParts()

	if ipb != "1234" {
		t.Errorf("unexpected result %v ", ipb)
	}
	if fpb != "0001" {
		t.Errorf("unexpected result %v ", fpb)
	}

	ipc, fpc := c.GetParts()

	if ipc != "-10" {
		t.Errorf("unexpected result %v ", ipc)
	}
	if fpc != "1" {
		t.Errorf("unexpected result %v ", fpc)
	}
}

func TestMultiplication(t *testing.T) {
	var a, b YapAmount
	var c YapCalculator

	as := "37310.09"
	bs := "0.1"
	if err := a.AmountFromString(as); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	if err := b.AmountFromString(bs); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	r, _ := c.Multiply(&a, &b)
	sr := r.AmountAsString()
	if s, err := strconv.ParseFloat(sr, 64); err == nil {
		if s != 3731.009 {
			t.Errorf("Multiplication failed  %v, %v, %v", as, bs, sr)
		}
	} else {
		t.Errorf("Conversion to float resulted in unexpected error %v", err)
	}

	as = "-37310.09"
	bs = "0.1"
	if err := a.AmountFromString(as); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	if err := b.AmountFromString(bs); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	r, _ = c.Multiply(&a, &b)
	sr = r.AmountAsString()
	if s, err := strconv.ParseFloat(sr, 64); err == nil {
		if s != -3731.009 {
			t.Errorf("Multiplication failed  %v, %v, %v", as, bs, sr)
		}
	} else {
		t.Errorf("Conversion to float resulted in unexpected error %v", err)
	}

	as = "37310.09"
	bs = "0"
	if err := a.AmountFromString(as); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	if err := b.AmountFromString(bs); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	r, _ = c.Multiply(&a, &b)
	if r.Value != 0 || r.Factor != 0 {
		t.Errorf("Multiplication failed  %v, %v, %v", as, bs, r)
	}

	as = "0"
	bs = "37310.09"
	if err := a.AmountFromString(as); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	if err := b.AmountFromString(bs); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	r, _ = c.Multiply(&a, &b)
	if r.Value != 0 || r.Factor != 0 {
		t.Errorf("Multiplication failed  %v, %v, %v", as, bs, r)
	}

	a.Value = 1 << 50
	a.Factor = 2
	b.Value = 1 << 50
	b.Factor = 1
	r, err := c.Multiply(&a, &b)
	if err == nil {
		t.Errorf("Overflow verification failed  %v, %v, %v", a, b, r)
	}

	a.Value = 1<<63 - 1
	a.Factor = 2
	b.Value = 1
	b.Factor = 1
	r, err = c.Multiply(&a, &b)
	if err != nil {
		t.Errorf("Overflow verification failed  %v, %v, %v", a, b, r)
	}

	a.Value = -(1 << 63)
	a.Factor = 2
	b.Value = 2
	b.Factor = 1
	r, err = c.Multiply(&a, &b)
	if err == nil {
		t.Errorf("Overflow verification failed  %v, %v, %v", a, b, r)
	}
}

func TestDivision(t *testing.T) {
	var a, b YapAmount
	var c YapCalculator

	as := "0.21"
	bs := "9.95"
	if err := a.AmountFromString(as); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	if err := b.AmountFromString(bs); err != nil {
		t.Errorf("Conversion resulted in unexpected error %v", err)
	}
	r, _ := c.Divide(&a, &b, 10)
	if r.Value != 211055276 || r.Factor != 10 {
		t.Errorf("Division failed  %v, %v, %v", as, bs, r)
	}

	a.Value = 2200
	a.Factor = 2
	b.Value = 70
	b.Factor = 1
	r, _ = c.Divide(&a, &b, 10)
	if r.Value != 31428571428 || r.Factor != 10 {
		t.Errorf("Division failed  %v, %v, %v", as, bs, r)
	}

	a.Value = 2200
	a.Factor = 2
	b.Value = 0
	b.Factor = 0
	_, err := c.Divide(&a, &b, 10)
	if err == nil {
		t.Errorf("Division by zero failed  %v, %v", a, b)
	}

	a.Value = 2200
	a.Factor = 2
	b.Value = 20
	b.Factor = 1
	r, _ = c.Divide(&a, &b, 10)
	if r.Value != 11 || r.Factor != 0 {
		t.Errorf("Division failed  %v, %v, %v", as, bs, r)
	}
}
