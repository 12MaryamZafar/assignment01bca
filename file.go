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

// NewBlock creates a new block with the provided transaction, nonce, and previous hash.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{ //new instance of the Block struct
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = block.CreateHash()
	return block
}

// CreateHash calculates and returns the SHA-256 hash of a block's data.
func (b *Block) CreateHash() string {
	data := b.Transaction + string(b.Nonce) + b.PreviousHash //concatenated string stored in data var.
	hashBytes := sha256.Sum256([]byte(data))                 //cryptographic hashing operation.
	return fmt.Sprintf("%x", hashBytes)                      //%x format specifier is used to represent the bytes as a hexadecimal string
}

// DisplayBlock prints the block's details.
func (b *Block) DisplayBlock() {
	fmt.Println("Transaction:", b.Transaction)
	fmt.Println("Nonce:", b.Nonce)
	fmt.Println("Previous Hash:", b.PreviousHash)
	fmt.Println("Current Hash:", b.CurrentHash)
}

// ChangeBlock updates the transaction of the given block.
func ChangeBlock(block *Block, newTransaction string) { //pointer to a Block struct
	block.Transaction = newTransaction
	block.CurrentHash = block.CreateHash()
}

// VerifyChain checks the integrity of the entire blockchain.
func VerifyChain(chain []*Block) bool { // a slice of pointers to Block structs.
	for i := 1; i < len(chain); i++ {
		currentBlock := chain[i]    //assigns the current block being checked
		previousBlock := chain[i-1] //assigns the previous block

		// Verify the current block's hash
		if currentBlock.CurrentHash != currentBlock.CreateHash() {
			return false
		}

		// Verify that the previous hash in the current block matches the hash of the previous block
		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}
	}
	return true
}

// CalculateHash calculates and returns the SHA-256 hash of a given string.
func CalculateHash(stringToHash string) string {
	hashBytes := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hashBytes)
}

func main() {
	// Create the genesis block
	genesisBlock := NewBlock("genesis", 0, "")

	// Create a new block
	block1 := NewBlock("bob to david", 123, genesisBlock.CurrentHash)

	// Create another new block
	block2 := NewBlock("alice to bob", 456, block1.CurrentHash)

	// Display all blocks in the blockchain
	genesisBlock.DisplayBlock()
	block1.DisplayBlock()
	block2.DisplayBlock()

	// Change the transaction of the second block
	ChangeBlock(block1, "alice to charlie")

	// Display the updated block
	fmt.Println("-------Updated Block-------")
	block1.DisplayBlock()

	// Verify the integrity of the blockchain
	isValid := VerifyChain([]*Block{genesisBlock, block1, block2})
	if isValid {
		fmt.Println("Blockchain is valid!")
	} else {
		fmt.Println("Blockchain is compromised!")
	}

	inputString := "Blockchain!"
	hash := CalculateHash(inputString)
	fmt.Println("Hash of input string:", hash)
}
