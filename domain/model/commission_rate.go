package model

const (
	defaultCommissionRate = 0.04
)

type CommissionRate float64

func (c CommissionRate) Value() float64 {
	return float64(c)
}

type NewCommissionRateInput struct{}

func NewCommissionRate(input *NewCommissionRateInput) (CommissionRate, error) {
	return CommissionRate(defaultCommissionRate), nil
}
