package model

const (
	defaultConsumptionTaxRate = 0.1
)

type ConsumptionTaxRate float64

func (c ConsumptionTaxRate) Value() float64 {
	return float64(c)
}

type NewConsumptionTaxRateInput struct{}

func NewConsumptionTaxRate(input *NewConsumptionTaxRateInput) (ConsumptionTaxRate, error) {
	return ConsumptionTaxRate(defaultConsumptionTaxRate), nil
}
