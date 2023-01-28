package blockchain

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
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
	data, err := db.Get([]byte("lastHash"), nil)
	if err == nil {
		bc.lashHash = Hash(data)
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
		log.Fatal("block can not be serialized")
	} else {
		err := bc.db.Put([]byte("b_"+b.hashCurr), bs, nil)
		if err != nil {
			log.Fatal("insert block error")

		}
	}
	bc.lashHash = b.hashCurr
	err := bc.db.Put([]byte("lastHash"), []byte(b.hashCurr), nil)
	if err != nil {
		log.Fatal("insert lasthash error")
	}
	return bc
}

func (bc *BlockChain) GetBlock(key Hash) (*Block, error) {
	bs, err := bc.db.Get([]byte("b_"+key), nil)
	if err != nil {
		log.Fatal("block cannot be retrieve", err)
	}
	b, err := BlockDeserialize(bs)
	if err != nil {
		log.Fatal("block cannot be deserialized", err)
	}
	return &b, err
}

// Clear only for test
func (bc *BlockChain) Clear() {
	bc.db.Delete([]byte("lastHash"), nil)
	iter := bc.db.NewIterator(util.BytesPrefix([]byte("b_")), nil)
	for iter.Next() {
		bc.db.Delete(iter.Key(), nil)
	}
	iter.Release()
	bc.lashHash = ""
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
