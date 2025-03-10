package consensus

import (
	"BlockChain_In_Go/core"
	"BlockChain_In_Go/utils"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math"
	"math/big"
	"time"
)

type PoofOfWork struct {
	block  *core.Block
	target *big.Int
}

// MiningDifficulty the target number of zeros in the hash
// const MiningDifficulty = 32
const MiningDifficulty = 16

// NewPoofOfWork init a target
func NewPoofOfWork(block *core.Block) *PoofOfWork {
	target := big.NewInt(1)
	// target = 0x1000...000(256 - 32 = 224 bits = 28 bytes)
	target.Lsh(target, uint(256-MiningDifficulty))
	pow := &PoofOfWork{block: block, target: target}
	return pow
}

func (pow *PoofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		utils.IntToHex(pow.block.TimeStamp),
		utils.IntToHex(int64(MiningDifficulty)),
		utils.IntToHex(int64(nonce)),
	}, []byte{})
	return data
}

func (pow *PoofOfWork) Run() (int, []byte) {
	var hash []byte
	var hashInt big.Int
	nonce := 0
	fmt.Printf("Mining Poof Of Work Data: %s\n", pow.block.Data)
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = crypto.Keccak256Hash(data).Bytes()
		hashInt.SetBytes(hash)
		// when hash less than target
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash
}

func (pow *PoofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := crypto.Keccak256Hash(data)
	hashInt.SetBytes(hash.Bytes())
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

func NewBlock(data []byte, prevBlock *core.Block) *core.Block {
	var block core.Block
	block.PrevBlockHash = prevBlock.Hash
	block.TimeStamp = time.Now().Unix()
	block.Data = data
	pow := NewPoofOfWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}

func NewGenesisBlock() *core.Block {
	return NewBlock([]byte("Genesis Block"), &core.Block{})
}

type PowBlockChain core.BlockChain

func (blockchain *PowBlockChain) AddBlock(data []byte) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, prevBlock)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func NewBlockChain() *PowBlockChain {
	blockchain := PowBlockChain{}
	genesisBlock := NewGenesisBlock()
	blockchain.Blocks = append(blockchain.Blocks, genesisBlock)
	return &blockchain
}
