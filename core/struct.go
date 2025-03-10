package core

type Block struct {
	// the hash of the previous block
	PrevBlockHash []byte
	// current time
	TimeStamp int64
	// store data,transaction...
	Data []byte
	// the hash of the current block
	Hash []byte
	// nonce
	Nonce int
}

type BlockChain struct {
	Blocks []*Block
}
