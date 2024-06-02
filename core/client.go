package core

type Client interface {
	Init()
	SetBC([]Block)
	GetBC([]Block)
	CreateTransaction() DefaultTransaction

	Validator
}
