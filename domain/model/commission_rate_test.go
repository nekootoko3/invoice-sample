package model

import "testing"

func TestNewCommissionRate(t *testing.T) {
	tests := []struct {
		name    string
		args    *NewCommissionRateInput
		want    CommissionRate
		wantErr bool
	}{
		{
			name:    "手数料率 0.04 を返却する",
			args:    nil,
			want:    CommissionRate(0.04),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewCommissionRate(tt.args); got != tt.want {
				t.Errorf("NewCommissionRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
