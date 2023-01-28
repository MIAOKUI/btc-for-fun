package blockchain

import (
	"fmt"
	"github.com/MIAOKUI/btc-for-fun/block"
	"github.com/MIAOKUI/btc-for-fun/pow"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"time"
)

type BlockChain struct {
	lashHash block.Hash
	db       *leveldb.DB
}

func NewBlockChain(db *leveldb.DB) *BlockChain {
	bc := &BlockChain{
		db: db,
	}
	data, err := db.Get([]byte("lastHash"), nil)
	if err == nil {
		bc.lashHash = block.Hash(data)
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
	b := block.NewBlock(bc.lashHash, txs)

	p := pow.NewPow(b)

	nonce, hash := p.Proof()
	if nonce == 0 || hash == "" {
		log.Fatal("block Hashcash Proof Failed")
	}
	b.SetNonce(nonce).SetHashCurr(hash)

	if bs, err := block.BlockSerialize(*b); err != nil {
		log.Fatal("block can not be serialized")
	} else {
		err := bc.db.Put([]byte("b_"+b.GetHashCurr()), bs, nil)
		if err != nil {
			log.Fatal("insert block error")

		}
	}
	bc.lashHash = b.GetHashCurr()
	err := bc.db.Put([]byte("lastHash"), []byte(b.GetHashCurr()), nil)
	if err != nil {
		log.Fatal("insert lasthash error")
	}
	return bc
}

func (bc *BlockChain) GetBlock(key block.Hash) (*block.Block, error) {
	bs, err := bc.db.Get([]byte("b_"+key), nil)
	if err != nil {
		log.Fatal("block cannot be retrieve", err)
	}
	b, err := block.BlockDeserialize(bs)
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
		fmt.Println("hashCurr: ", b.GetHashCurr())
		fmt.Println("txs: ", b.GetTxs())
		fmt.Println("time: ", b.GetHeader().GetTime().Format(time.Layout))
		fmt.Println("HashPrevious: ", b.GetHeader().GetHashPrevBlock())
		curHash = b.GetHeader().GetHashPrevBlock()
		fmt.Println()
	}
}
