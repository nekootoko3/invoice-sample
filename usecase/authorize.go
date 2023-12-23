package usecase

import "github.com/nekootoko3/invoice-sample/domain/model"

type AuthorizeActor struct {
	User    *model.User
	Company *model.Company
}

type AuthorizeResource struct {
	Invoice *model.Invoice
}

type AuthorizeAction int

const (
	AuthorizeActionCreate AuthorizeAction = iota + 1
	AuthorizeActionRead
	AuthorizeActionUpdate
	AuthorizeActionDelete
)

func Authorize(actor AuthorizeActor, resource AuthorizeResource, action AuthorizeAction) bool {
	if actor.User == nil || actor.Company == nil {
		return false
	}

	if actor.User.CompanyID != actor.Company.ID {
		return false
	}

	return true
}
