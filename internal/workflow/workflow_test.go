package workflow

import (
	"testing"

	"github.com/mmuoDev/price-discount/internal"
	"github.com/stretchr/testify/assert"
	
)

func Test30PercentDiscountIsAppliedAsExpected(t *testing.T) {
	p := internal.Product{
		Name: "LV Shoes",
		SKU: "00844",
		Category: "boots",
		Price: 125,
	}
	res := computeDiscount(p)
	t.Run("New price is as expected", func(t *testing.T) {
		assert.Equal(t, int64(88), res.NewAmount)
	})
	t.Run("Discount is as expected", func(t *testing.T) {
		assert.Equal(t, "30", res.DiscountedPercentage)
	})
}

func Test15PercentDiscountIsAppliedAsExpected(t *testing.T) {
	p := internal.Product{
		Name: "LV Bags",
		SKU: "000003",
		Category: "bags",
		Price: 125,
	}
	res := computeDiscount(p)
	t.Run("New price is as expected", func(t *testing.T) {
		assert.Equal(t, int64(107), res.NewAmount)
	})
	t.Run("Discount is as expected", func(t *testing.T) {
		assert.Equal(t, "15", res.DiscountedPercentage)
	})
}

func TestMultipleDiscountsIsAppliedAsExpected(t *testing.T) {
	p := internal.Product{
		Name: "LV Bags",
		SKU: "000003",
		Category: "boots",
		Price: 125,
	}
	res := computeDiscount(p)
	t.Run("New price is as expected", func(t *testing.T) {
		assert.Equal(t, int64(88), res.NewAmount)
	})
	t.Run("Discount is as expected", func(t *testing.T) {
		assert.Equal(t, "30", res.DiscountedPercentage)
	})
}