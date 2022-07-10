package main

// import (
// 	"fmt"

// 	"github.com/namle133/go-blockchain.git/go-blockchain/blockchain"
// 	// "strconv"
// )

// func main() {
// 	// fmt.Println(string([]byte{0}))
// 	// data := []byte("hello")
// 	// hash := sha256.Sum256(data)
// 	// fmt.Printf("%x", hash[:])

// 	chain := blockchain.InitBlockChain()

// 	chain.AddBlock("First Block after Genesis")
// 	chain.AddBlock("Second Block after Genesis")
// 	chain.AddBlock("Third Block after Genesis")

// 	for _, block := range chain.Blocks {
// 		fmt.Printf("Previous Hash: %x\n", block.PreHash)
// 		fmt.Printf("Data in Block: %s\n", block.Data)
// 		fmt.Printf("Hash: %x\n", block.Hash)
// 		// fmt.Println(strconv.Itoa(block.Hash))
// 		// fmt.Println(reflect.TypeOf(block.Hash))
// 		// fmt.Println(block.Hash)
// 		fmt.Printf("Nonce: %v\n", block.Nonce)
// 		fmt.Printf("Time: %s\n", block.TimeStamp)

// 	}
// 	fmt.Printf("valid: %v\n\n", chain.IsValid())
// 	chain.Blocks[1].Data = []byte("Namle")
// 	chain.Blocks[1].Hash = chain.Blocks[1].CalculateHash()
// 	// for _, block := range chain.Blocks {
// 	// 	fmt.Printf("Previous Hash: %x\n", block.PreHash)
// 	// 	fmt.Printf("Data in Block: %s\n", block.Data)
// 	// 	fmt.Printf("Hash: %x\n", block.Hash)
// 	// }
// 	fmt.Printf("valid: %v\n\n", chain.IsValid())

// }

import (
	"fmt"
	"strconv"

	"github.com/namle133/go-blockchain.git/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {

		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
	//Testing the Validate
	chain.Blocks[1].Data = []byte("Nam")
	pow := blockchain.NewProof(chain.Blocks[1])
	_, k := pow.Run()
	chain.Blocks[1].Hash = k
	for _, block := range chain.Blocks {

		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}

}
