package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Block struct {
	Hash      []byte
	TimeStamp []byte
	Data      []byte
	PreHash   []byte
	Nonce     int
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) DeriveHash() {
	times := time.Now().String()
	info := bytes.Join([][]byte{[]byte(times), b.Data, b.PreHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
	b.TimeStamp = []byte(times)
}

// func (b *Block) GiveHash(h []byte) {
// 	b.Hash = h[:]
// }

func (b *Block) CalculateHash() []byte {
	info := bytes.Join([][]byte{b.TimeStamp, b.Data, b.PreHash}, []byte{})
	hash := sha256.Sum256(info)
	h := hash[:]
	return h
}

func CreatBlock(data string, preHash []byte) *Block {

	block := &Block{[]byte{}, []byte{}, []byte(data), preHash, 0}
	// h := Hash1(bytes.Join([][]byte{block.TimeStamp, block.Data, block.PreHash}, []byte{}))
	// block.GiveHash(h)
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreatBlock(data, prevBlock.Hash)
	new.Mine()
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreatBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func ComparePassword(hashedPassword []byte, password []byte) error {
	er := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if er != nil {
		return er
	}
	return nil
}

func Hash1(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func (b *Block) Mine() {
	// h := Hash1(b.Hash)
	for string(b.Hash[0:1]) != "0" {
		// fmt.Println([]byte("0"))
		// fmt.Println(string(h[0:1]))
		b.Nonce++
		// h = Hash1(h)
		b.DeriveHash()
	}
	// b.GiveHash(h)
}

func (chain *BlockChain) IsValid() bool {
	for i := 1; i < len(chain.Blocks); i++ {
		currentBlock := chain.Blocks[i]
		preBlock := chain.Blocks[i-1]
		if string(currentBlock.Hash) != string(currentBlock.CalculateHash()) {
			return false
		}

		if string(currentBlock.PreHash) != string(preBlock.Hash) {
			return false
		}
	}
	return true
}
