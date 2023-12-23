package model

import (
	"context"

	"github.com/nekootoko3/invoice-sample/util/timeutil"
	"github.com/oklog/ulid/v2"
)

type InvoiceStatus int

const (
	InvoiceStatusUnprocessed InvoiceStatus = iota + 1
	InvoiceStatusProcessing
	InvoiceStatusPaid
	InvoiceStatusError
)

type Invoice struct {
	ID                 string
	CompanyID          string
	ClientID           string
	IssueDate          string
	DueDate            string
	PaymentAmount      float64
	Commission         float64
	CommissionRate     CommissionRate
	ConsumptionTax     float64
	ConsumptionTaxRate ConsumptionTaxRate
	InvoiceAmount      float64
	Status             InvoiceStatus
}

type NewInvoiceInput struct {
	ID                      string // TODO: テストのために id を受け取るようにしているが mock を挟めるようにする
	CompanyID               string
	ClientID                string
	IssueDate               string
	DueDate                 string
	PaymentAmount           float64
	CommissionRateInput     CommissionRate
	ConsumptionTaxRateInput ConsumptionTaxRate
}

func NewInvoice(
	input NewInvoiceInput,
) (*Invoice, error) {
	issueDate, err := timeutil.ParseDate(input.IssueDate)
	if err != nil {
		return nil, err
	}
	dueDate, err := timeutil.ParseDate(input.DueDate)
	if err != nil {
		return nil, err
	}
	if !dueDate.After(issueDate) {
		return nil, ErrInvalidDueDate
	}

	commission := input.PaymentAmount * input.CommissionRateInput.Value()
	consumptionTax := input.PaymentAmount * input.ConsumptionTaxRateInput.Value()
	invoiceAmount := input.PaymentAmount + commission + consumptionTax

	id := input.ID
	if id == "" {
		id = ulid.Make().String()
	}

	return &Invoice{
		ID:                 id,
		CompanyID:          input.CompanyID,
		ClientID:           input.ClientID,
		IssueDate:          input.IssueDate,
		DueDate:            input.DueDate,
		PaymentAmount:      input.PaymentAmount,
		CommissionRate:     input.CommissionRateInput,
		Commission:         commission,
		ConsumptionTaxRate: input.ConsumptionTaxRateInput,
		ConsumptionTax:     consumptionTax,
		InvoiceAmount:      invoiceAmount,
		Status:             InvoiceStatusUnprocessed,
	}, nil
}

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, invoice Invoice) error
}
