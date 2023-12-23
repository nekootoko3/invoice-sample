package usecase

import (
	"context"

	"github.com/nekootoko3/invoice-sample/domain/model"
)

type CreateInvoiceInput struct {
	User    *model.User
	Company *model.Company
	Client  *model.Client

	IssueDate     string
	PaymentAmount float64
	DueDate       string
}

type CreateInvoice struct {
	invoiceRepository model.InvoiceRepository
}

func (c *CreateInvoice) Execute(ctx context.Context, input *CreateInvoiceInput) error {
	if !Authorize(
		AuthorizeActor{User: input.User, Company: input.Company},
		AuthorizeResource{&model.Invoice{}},
		AuthorizeActionCreate,
	) {
		return ErrUnauthorized
	}

	invoice, err := model.NewInvoice(model.NewInvoiceInput{
		CompanyID:     input.Company.ID,
		ClientID:      input.Client.ID,
		IssueDate:     input.IssueDate,
		PaymentAmount: input.PaymentAmount,
		DueDate:       input.DueDate,
	})
	if err != nil {
		return err
	}

	return c.invoiceRepository.Create(ctx, invoice)
}
