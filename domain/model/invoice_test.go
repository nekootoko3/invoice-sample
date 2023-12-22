package model

import (
	"reflect"
	"testing"
)

func TestNewInvoice(t *testing.T) {
	type args struct {
		id                 string
		issueDate          string
		dueDate            string
		paymentAmount      float64
		commissionRate     float64
		consumptionTaxRate float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Invoice
		wantErr bool
	}{
		{
			name: "発行日より支払日が未来の場合、請求を返却する",
			args: args{
				id:                 "test",
				issueDate:          "2020-01-01",
				dueDate:            "2020-01-31",
				paymentAmount:      10000,
				commissionRate:     0.1,
				consumptionTaxRate: 0.1,
			},
			want: &Invoice{
				ID:                 "test",
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
			args: args{
				id:                 "test",
				issueDate:          "2020-01-01",
				dueDate:            "2020-01-01",
				paymentAmount:      10000,
				commissionRate:     0.1,
				consumptionTaxRate: 0.1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInvoice(tt.args.id, tt.args.issueDate, tt.args.dueDate, tt.args.paymentAmount, tt.args.commissionRate, tt.args.consumptionTaxRate)
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
