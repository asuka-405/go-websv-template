package libblockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"
)

type BlockLite struct {
	index     int
	timestamp string
	bpm       int
	hash      string
	prevHash  string
}

func (block *BlockLite) GetIndex() int {
	return block.index
}

func (block *BlockLite) GetTimestamp() string {
	return block.timestamp
}

func (block *BlockLite) GetBPM() int {
	return block.bpm
}

func (block *BlockLite) GetHash() string {
	return block.hash
}

func (block *BlockLite) GetPrevHash() string {
	return block.prevHash
}

func (BlockLite) New(BPM int, prev *BlockLite) (BlockLite, error) {
	index := 0
	var prevHash string
	curTime := time.Now().String()
	if prev != nil {
		index = prev.index + 1
		prevHash = prev.hash
	}
	record := strconv.Itoa(index) + curTime + strconv.Itoa(BPM) + prevHash
	return BlockLite{
		index:     index,
		timestamp: curTime,
		bpm:       BPM,
		hash:      HashString(record),
		prevHash:  prevHash,
	}, nil
}

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (block *BlockLite) CalculateHash() string {
	record := strconv.Itoa(block.index) + block.timestamp + strconv.Itoa(block.bpm) + block.prevHash
	return HashString(record)
}

func (block *BlockLite) isValid(prev *BlockLite) bool {
	if prev == nil {
		return true
	}
	if prev.index+1 != block.index {
		fmt.Print("index err: ")
		return false
	} else if prev.hash != block.prevHash {
		fmt.Print("prevhash err: ")
		return false
	} else if block.CalculateHash() != block.hash {
		fmt.Print("hash mismatch err: \n")
		fmt.Print("given hash: " + block.hash + "\n")
		fmt.Print("calc  hash: " + block.CalculateHash() + "\n")
		return false
	}
	return true
}

func (block *BlockLite) Render() string {
	return fmt.Sprintf(`
		<div class="blockchain-block">
			<div class="block-id">ID: %d</div>
			<div class="block-data">DATA: %d</div>
			<div class="block-hash">HASH: %s</div>
			<div class="block-prevhash">PREVHASH: %s</div>
			<div class="block-timestamp">TIMESTAMP: %s</div>
		</div>`, block.index, block.bpm, block.hash, block.prevHash, block.timestamp)
}

type BlockChainLite struct {
	Chain  []BlockLite
	Head   *BlockLite
	Tail   *BlockLite
	Length int
}

func (BlockChainLite) New() BlockChainLite {
	var chain BlockChainLite
	block, _ := BlockLite{}.New(80, nil)
	chain.Head = &block
	chain.Tail = &block
	chain.Length = 1
	chain.Chain = append(chain.Chain, block)
	return chain
}

func (blockchain *BlockChainLite) Push(block *BlockLite) {
	if !block.isValid(blockchain.Tail) {
		log.Fatal("cannot push invalid block, skipping: " + fmt.Sprint(block))
	}
	blockchain.Tail = block
	blockchain.Chain = append(blockchain.Chain, *block)
	blockchain.Length++
}

func (blockchain *BlockChainLite) getTrueChain(newChain, oldChain BlockChainLite) BlockChainLite {
	if len(newChain.Chain) > len(oldChain.Chain) {
		return newChain
	}
	return oldChain
}

func (blockchain *BlockChainLite) Render() string {
	html := ``

	for _, block := range blockchain.Chain {
		html += block.Render()
	}

	finalHtml := fmt.Sprintf(`<div class="blockchain-container">
								<svg id="arrowCanvas"></svg>
								%s
							  </div>`, html)
	return finalHtml
}
