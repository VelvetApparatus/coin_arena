package core

type Validator interface {
	CheckSequence(transaction *DefaultTransaction) bool
	BC() []Block
	Conditions() []Condition
}

type Condition interface {
	Valid() bool
}
