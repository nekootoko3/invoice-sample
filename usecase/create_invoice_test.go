package usecase

import (
	"context"
	"testing"

	"github.com/nekootoko3/invoice-sample/domain/model"
)

func TestCreateInvoice_Execute(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		invoiceMock := &model.InvoiceRepositoryMock{
			CreateFunc: func(ctx context.Context, invoice *model.Invoice) error {
				return nil
			},
		}

		usecase := &CreateInvoice{
			invoiceRepository: invoiceMock,
		}

		input := &CreateInvoiceInput{
			User: &model.User{
				CompanyID: "company-1",
			},
			Company: &model.Company{
				ID: "company-1",
			},
			Client: &model.Client{
				ID: "client-1",
			},
			IssueDate:     "2020-01-01",
			DueDate:       "2020-01-31",
			PaymentAmount: 1000,
		}

		err := usecase.Execute(context.Background(), input)
		if err != nil {
			t.Error("error should be nil")
		}
		if len(invoiceMock.CreateCalls()) != 1 {
			t.Error("Create should be called once")
		}
	})
}
