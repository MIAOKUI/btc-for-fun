package main

import (
	"github.com/MIAOKUI/btc-for-fun/blockchain"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	//bc := blockchain.NewBlock("", "Gensis Block.")
	//fmt.Print(bc)
	dbpath := "testdb"
	db, err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		log.Fatal(err)
	}

	bc := blockchain.NewBlockChain(db)
	bc.AddGenesisBlock()
	bc.AddBlock("first").AddBlock("second").AddBlock("third")
	bc.PrintIterate()
	//fmt.Println(bc)
}
