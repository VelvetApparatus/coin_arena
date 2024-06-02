package core

import "math/big"

type Transaction interface {
	From() string
	To() string
	Amount() *big.Float
}

type DefaultTransaction struct {
	from   string
	to     string
	amount *big.Float
}

func NewTransaction(from, to string, amount *big.Float) *DefaultTransaction {
	return &DefaultTransaction{
		from:   from,
		to:     to,
		amount: amount,
	}
}

func (t *DefaultTransaction) From() string { return t.from }

func (t *DefaultTransaction) To() string { return t.to }

func (t *DefaultTransaction) Amount() *big.Float { return t.amount }
