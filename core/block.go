package core

type Block interface {
	Transactions() []DefaultTransaction
	PreviousHash() []byte
}
