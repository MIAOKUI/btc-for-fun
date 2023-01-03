package main

import (
	"fmt"
	"github.com/MIAOKUI/btc-for-fun/blockchain"
)

func main() {
	b := blockchain.NewBlock("", "Gensis Block.")
	fmt.Print(b)
}
