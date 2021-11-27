package core

// define a block.
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// chain is a just slice of blocks
type BlockChain struct {
	Blocks []*Block
}
