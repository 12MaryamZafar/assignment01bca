package main

import (
	"crypto/sha256"
	"fmt"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// Blockchain represents a chain of blocks.
type Blockchain struct {
	Blocks []Block
}

// NewBlock creates a new block with the provided transaction, nonce, and previous hash.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = block.CreateHash()
	return block
}

// CreateHash calculates and returns the SHA-256 hash of a block's data.
func (b *Block) CreateHash() string {
	data := b.Transaction + string(b.Nonce) + b.PreviousHash
	hashBytes := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hashBytes)
}

// DisplayBlocks prints all the blocks in a nice format.
func (bc *Blockchain) DisplayBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i+1)
		fmt.Printf("  Transaction: %s\n", block.Transaction)
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Current Hash: %s\n", block.CurrentHash)
	}
}

func main() {
	// Create a new blockchain
	blockchain := Blockchain{}

	// Add some blocks to the blockchain
	blockchain.Blocks = append(blockchain.Blocks, *NewBlock("bob to alice", 123, "genesisHash"))
	blockchain.Blocks = append(blockchain.Blocks, *NewBlock("alice to bob", 456, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash))
	blockchain.Blocks = append(blockchain.Blocks, *NewBlock("charlie to david", 789, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash))

	// Display all the blocks
	blockchain.DisplayBlocks()
}
