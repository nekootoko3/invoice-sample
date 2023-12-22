package model

import "time"

type InvoiceData struct {
	ID                 string
	IssueDate          time.Time
	PaymentAmount      float64
	Commission         float64
	CommissionRate     float64
	ConsumptionTax     float64
	ConsumptionTaxRate float64
	InvoiceAmount      float64
	DueDate            time.Time
	Status             string
}
