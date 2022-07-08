package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DerviHash() {
	info := bytes.Join([][]byte{b.Data, b.PreHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreatBlock(data string, preHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), preHash}
	block.DerviHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreatBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreatBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)

	}
}
