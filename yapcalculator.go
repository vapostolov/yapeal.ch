package yapstones

// YapCalculator contains calculators used in amount and money
type YapCalculator struct{}

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
