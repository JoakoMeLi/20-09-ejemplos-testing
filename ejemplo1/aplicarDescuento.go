package discount

type DiscountCalculator struct {
	MinimumPurchaseAmount int
	DiscountAmount        int
}

func NewDiscountCalculator(MinimumPurchaseAmount int, DiscountAmount int) *DiscountCalculator {
	return &DiscountCalculator{
		MinimumPurchaseAmount: MinimumPurchaseAmount,
		DiscountAmount:        DiscountAmount,
	}
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.MinimumPurchaseAmount {
		return purchaseAmount - c.DiscountAmount
	}
	return purchaseAmount
}
