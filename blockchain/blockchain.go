package blockchain

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"time"
)

type BlockChain struct {
	lashHash Hash
	db       *leveldb.DB
}

func NewBlockChain(db *leveldb.DB) *BlockChain {
	bc := &BlockChain{
		db: db,
	}
	return bc
}

func (bc *BlockChain) AddGenesisBlock() *BlockChain {
	if bc.lashHash != "" {
		return bc
	}
	return bc.AddBlock("The Genesis Block")
}

func (bc *BlockChain) AddBlock(txs string) *BlockChain {
	b := NewBlock(bc.lashHash, txs)
	if bs, err := BlockSerialize(*b); err != nil {
		bc.db.Put([]byte(b.hashCurr), bs, nil)
	} else {
		log.Fatal("block can not be serialized")
	}
	bc.lashHash = b.hashCurr
	return bc
}

func (bc *BlockChain) GetBlock(key Hash) (*Block, error) {
	bs, err := bc.db.Get([]byte(key), nil)
	if err != nil {
		log.Fatal("block cannot be retrieve")
	}
	b, err := BlockDeserialize(bs)
	if err != nil {
		log.Fatal("block cannot be deserialized")
	}
	return b, err
}

func (bc *BlockChain) PrintIterate() {
	for curHash := bc.lashHash; curHash != ""; {
		b, err := bc.GetBlock(curHash)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hashCurr: ", b.hashCurr)
		fmt.Println("txs: ", b.txs)
		fmt.Println("time: ", b.header.time.Format(time.Layout))
		fmt.Println("HashPrevious: ", b.header.hashPrevBlock)
		curHash = b.header.hashPrevBlock
		fmt.Println()
	}
}
