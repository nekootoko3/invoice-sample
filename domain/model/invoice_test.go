package model

import (
	"reflect"
	"testing"
)

func TestNewInvoice(t *testing.T) {
	tests := []struct {
		name    string
		args    NewInvoiceInput
		want    *Invoice
		wantErr bool
	}{
		{
			name: "発行日より支払日が未来の場合、請求を返却する",
			args: NewInvoiceInput{
				ID:                      "test",
				IssueDate:               "2020-01-01",
				DueDate:                 "2020-01-31",
				PaymentAmount:           10000,
				CommissionRateInput:     0.1,
				ConsumptionTaxRateInput: 0.1,
				CompanyID:               "company-1",
				ClientID:                "client-1",
			},
			want: &Invoice{
				ID:                 "test",
				CompanyID:          "company-1",
				ClientID:           "client-1",
				IssueDate:          "2020-01-01",
				DueDate:            "2020-01-31",
				PaymentAmount:      10000,
				Commission:         1000,
				CommissionRate:     0.1,
				ConsumptionTax:     1000,
				ConsumptionTaxRate: 0.1,
				InvoiceAmount:      12000,
				Status:             InvoiceStatusUnprocessed,
			},
			wantErr: false,
		},
		{
			name: "発行日より支払日が未来ではない場合、エラーを返却する",
			args: NewInvoiceInput{
				ID:        "test",
				IssueDate: "2020-01-01",
				DueDate:   "2020-01-01",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInvoice(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
