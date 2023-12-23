package usecase

import (
	"testing"

	"github.com/nekootoko3/invoice-sample/domain/model"
)

func TestAuthorize(t *testing.T) {
	type args struct {
		actor    AuthorizeActor
		resource AuthorizeResource
		action   AuthorizeAction
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "actor の User.CompanyID と Company.ID が一致する場合、trueを返す",
			args: args{
				actor: AuthorizeActor{
					User:    &model.User{CompanyID: "company-1"},
					Company: &model.Company{ID: "company-1"},
				},
				resource: AuthorizeResource{},
				action:   AuthorizeActionCreate,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Authorize(tt.args.actor, tt.args.resource, tt.args.action); got != tt.want {
				t.Errorf("Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}
