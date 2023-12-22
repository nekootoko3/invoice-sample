package model

type Client struct {
	ID             string
	CompanyName    string
	Representative string
	PhoneNumber    string
	ZipCode        string
	Address        string
	BankAccount    ClientBankAccount
}
