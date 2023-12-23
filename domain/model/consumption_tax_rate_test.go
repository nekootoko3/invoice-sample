package model

import "testing"

func TestNewConsumptionTaxRate(t *testing.T) {
	tests := []struct {
		name    string
		args    *NewConsumptionTaxRateInput
		want    ConsumptionTaxRate
		wantErr bool
	}{
		{
			name:    "消費税率 0.1 を返却する",
			args:    nil,
			want:    ConsumptionTaxRate(0.1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewConsumptionTaxRate(tt.args); got != tt.want {
				t.Errorf("NewConsumptionTaxRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
