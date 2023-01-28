package main

import (
	"flag"
	"fmt"
	"github.com/MIAOKUI/btc-for-fun/blockchain"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"strings"
)

func main() {
	bc, err := InitBlockChain()
	if err != nil {
		log.Fatal(err)
	}

	arg1 := ""
	if len(os.Args) >= 2 {
		arg1 = os.Args[1]
	}
	switch strings.ToLower(arg1) {
	case "help":
		fallthrough
	case "create:block":
		fs := flag.NewFlagSet("create:block", flag.ExitOnError)
		txs := fs.String("txs", "", "")
		fs.Parse(os.Args[2:])
		bc.AddBlock(*txs)
	case "show":
		bc.PrintIterate()
	case "init":
		bc.Clear()
	default:
		Usage()
	}

}

func InitBlockChain() (*blockchain.BlockChain, error) {
	dbpath := "db"
	db, err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		log.Fatal(err)
	}

	bc := blockchain.NewBlockChain(db)
	bc.AddGenesisBlock()
	return bc, nil
}

func Usage() {
	fmt.Println("bcli is cmd line tool for BTC control")
	fmt.Println()
	fmt.Println("Usage")
	fmt.Printf("\t%s\t\t%s\n", "bcli init", "inital blockchain")
	fmt.Printf("\t%s\t\t%s\n", "bcli create:block <txs>", "create block on blockchain")
	fmt.Printf("\t%s\t\t%s\n", "bcli help", "help info")
	fmt.Printf("\t%s\t\t%s\n", "bcli show", "show all blocks")
}
