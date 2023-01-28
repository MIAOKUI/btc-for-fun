package block

import (
	"fmt"
	"time"
)

type Hash = string

const HashLen = 256
const nodeVersion = 0
const blockBits = 12

type Block struct {
	header    Header
	txs       string
	txCounter int
	hashCurr  Hash
}

type Header struct {
	version        int
	hashPrevBlock  Hash
	hashMerkleRoot Hash
	time           time.Time
	bits           int
	nonce          int
}

func NewBlock(prevHash Hash, txs string) *Block {
	b := &Block{
		header: Header{
			version:       nodeVersion,
			hashPrevBlock: prevHash,
			time:          time.Now(),
			bits:          blockBits,
		},
		txs:       txs,
		txCounter: 1,
		hashCurr:  "",
	}
	return b
}

func (bh *Header) Stringify() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.hashPrevBlock,
		bh.hashMerkleRoot,
		bh.time.UnixNano(),
		bh.bits,
		bh.nonce)
}

func (bh *Header) GetTime() time.Time {
	return bh.time
}

func (bh *Header) GetHashPrevBlock() Hash {
	return bh.hashPrevBlock
}

func (b *Block) GetServiceString() string {
	return fmt.Sprintf("%d%s%s%d%s%d",
		b.header.version,
		b.header.hashPrevBlock,
		b.header.hashMerkleRoot,
		b.header.time.String(),
		b.header.bits)
}

func (b *Block) SetHashCurr(hash Hash) *Block {
	b.hashCurr = hash
	return b
}

func (b *Block) GetHashCurr() Hash {
	return b.hashCurr
}

func (b *Block) GetTxs() string {
	return b.txs
}

func (b *Block) GetHeader() *Header {
	return &b.header
}

func (b *Block) SetNonce(nonce int) *Block {
	b.header.nonce = nonce
	return b
}

func (b *Block) GetBits() int {
	return b.header.bits
}
