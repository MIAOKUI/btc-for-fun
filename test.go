package main

import (
	"fmt"
	"github.com/MIAOKUI/btc-for-fun/blockchain"
)

func main() {
	//b := blockchain.NewBlock("", "Gensis Block.")
	//fmt.Print(b)

	bc := blockchain.NewBlockChain()
	bc.AddGenesisBlock()
	bc.AddBlock("first").AddBlock("second")
	fmt.Println(bc)
}
