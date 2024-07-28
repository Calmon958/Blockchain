package main

import (
"crypto/sha256"
"encoding/json"
"fmt"
"time"
"strings"
"log"
)

type Block struct {
nonce int
previousHash string
timestamp int64
transactions []string
}

//First Block in the blockchain.
func NewBlock(nonce int, previousHash string) *Block {
b := new(Block)
b.timestamp = time.Now().UnixNano()
b.nonce = nonce
b.previousHash = previousHash
return b

/* implimentation of the same kind of code as above.
return &Block{
timestamp: time.Now().UnixNano(),
}
*/
}

//Structural output for displaying the blockchain.
func (b *Block) Print() {
fmt.Printf("timestamp         %d\n", b.timestamp)
fmt.Printf("nonce              %d\n", b.nonce)
fmt.Printf("Previous Hash      %s\n", b.previousHash)
fmt.Printf("transactions       %s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
m, _ := json.Marshal(b)
fmt.Println(string(m))
return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
return json.Marshal(struct {
Timestamp int64 `json:"timestamp"`
Nonce int  `json:"nonce"'
PreviousHash string  `json:"previousHash"`
Transactions []string `json:"transactions"`
}{
Timestamp: b.timestamp,
Nonce: b.nonce,
PreviousHash: b.previousHash,
Transactions: b.transactions,
})
}


func init() {
log.SetPrefix("Blockchain: ")
}

//create structure for our blockchain
type Blockchain struct {
transactionPool []string
chain  []*Block
}


func NewBlockchain() *Blockchain {
bc:=new(Blockchain)
bc.CreateBlock(0, "init hash")
return bc
}

func (bc *Blockchain)  CreateBlock(nonce int, prevHash string) *Block {

b := NewBlock(nonce, prevHash)
bc.chain = append(bc.chain, b)
return b
}

func (bc *Blockchain) Print() {
for  i, block := range bc.chain{
fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
block.Print()
}
fmt.Printf("%s\n", strings.Repeat("*", 25))
}


func main(){
/*
b:=NewBlock(0, "initial hash value")
//fmt.Println(b)
b.Print()
*/

block := &Block{nonce: 1}
fmt.Printf("%x\n", block.Hash())


/*
blockChain := NewBlockchain()
blockChain.Print()
fmt.Println()
blockChain.CreateBlock(5, "hash 1")
blockChain.Print()
fmt.Println()
blockChain.CreateBlock(3, "hash 2")
blockChain.Print()
*/

//fmt.Println(blockChain)
}

