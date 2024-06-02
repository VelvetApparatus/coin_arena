package core

type GenesisBlock interface {
	GenerateGenesisBlock() Block
}
