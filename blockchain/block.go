package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Hash = string

const node_version = 0

type Block struct {
	header    BlockHeader
	txs       string
	txCounter int
	hashCurr  Hash
}

type BlockHeader struct {
	version        int
	hashPrevBlock  Hash
	hashMerkleRoot Hash
	time           time.Time
	bits           int
	nonce          int
}

func NewBlock(prevHash Hash, txs string) *Block {
	b := &Block{
		header: BlockHeader{
			version:       node_version,
			hashPrevBlock: prevHash,
			time:          time.Now(),
		},
		txs:       txs,
		txCounter: 1,
		hashCurr:  "",
	}
	b.SetHashCurr()
	return b
}

func (bh *BlockHeader) Stringify() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.hashPrevBlock,
		bh.hashMerkleRoot,
		bh.time.UnixNano(),
		bh.bits,
		bh.nonce)
}

func (b *Block) GetServiceString() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		b.header.version,
		b.header.hashPrevBlock,
		b.header.hashMerkleRoot,
		b.header.time,
		b.header.bits)
}

func (b *Block) SetHashCurr() *Block {
	headerStr := b.header.Stringify()
	b.hashCurr = fmt.Sprintf("%x", sha256.Sum256([]byte(headerStr)))
	return b
}

func (b *Block) GetBits() int {
	return b.header.bits
}
