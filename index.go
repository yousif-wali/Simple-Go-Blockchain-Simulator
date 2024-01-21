package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Timestamp    string
	Data         string
	PrevBlockHash string
	Hash         string
}

// Blockchain is a series of validated Blocks
type Blockchain struct {
	Head *Block
	Tail *Block
}

// NewBlock creates a new Block using previous block's hash
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{time.Now().Format("2006-01-02 15:04:05"), data, prevBlockHash, ""}
	block.Hash = block.calculateHash()
	return block
}

// AddBlock adds a new block to the Blockchain
func (bc *Blockchain) AddBlock(data string) {
	newBlock := NewBlock(data, bc.Tail.Hash)
	if bc.Head == nil {
		bc.Head = newBlock
		bc.Tail = newBlock
	} else {
		bc.Tail = newBlock
	}
}

// calculateHash calculates the SHA256 hash of a Block
func (b *Block) calculateHash() string {
	record := b.Timestamp + b.Data + b.PrevBlockHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// NewBlockchain creates a new Blockchain with an initial block
func NewBlockchain() *Blockchain {
	return &Blockchain{NewBlock("Initial Block", ""), nil}
}

// PrintChain prints the blocks in the blockchain
func (bc *Blockchain) PrintChain() {
	for block := bc.Head; block != nil; block = block {
		fmt.Printf("Prev. hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Block #1")
	bc.AddBlock("Block #2")

	bc.PrintChain()
}
