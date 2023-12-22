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
	IssueDate          string
	DueDate            string
	PaymentAmount      float64
	Commission         float64
	CommissionRate     float64
	ConsumptionTax     float64
	ConsumptionTaxRate float64
	InvoiceAmount      float64
	Status             InvoiceStatus
}

func NewInvoice(
	Id string, // TODO: テストのために id を受け取るようにしているが mock を挟めるようにする
	IssueDate string,
	DueDate string,
	PaymentAmount float64,
	CommissionRate float64,
	ConsumptionTaxRate float64,
) (*Invoice, error) {
	issueDate, err := timeutil.ParseDate(IssueDate)
	if err != nil {
		return nil, err
	}
	dueDate, err := timeutil.ParseDate(DueDate)
	if err != nil {
		return nil, err
	}
	if !dueDate.After(issueDate) {
		return nil, ErrInvalidDueDate
	}

	commission := PaymentAmount * CommissionRate
	consumptionTax := PaymentAmount * ConsumptionTaxRate
	invoiceAmount := PaymentAmount + commission + consumptionTax

	id := Id
	if id == "" {
		id = ulid.Make().String()
	}

	return &Invoice{
		ID:                 id,
		IssueDate:          IssueDate,
		DueDate:            DueDate,
		PaymentAmount:      PaymentAmount,
		CommissionRate:     CommissionRate,
		Commission:         commission,
		ConsumptionTaxRate: ConsumptionTaxRate,
		ConsumptionTax:     consumptionTax,
		InvoiceAmount:      invoiceAmount,
		Status:             InvoiceStatusUnprocessed,
	}, nil
}

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, invoice Invoice) error
}
