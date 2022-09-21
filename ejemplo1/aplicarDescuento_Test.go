package discount_test

import (
	"testing"

	discount "github.com/JoakoMeLi/20-09-ejemplos-testing/ejemplo1"
)

func TestDiscountApplied(t *testing.T) {
	calculator := discount.NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)

	if 130 != amount {
		t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := discount.NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(50)

	if 50 != amount {
		t.Fail()
	}
}
