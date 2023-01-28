package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

type Block struct {
	CurrHash string
	TXs      string
}

func main() {
	bits := 20
	target := big.NewInt(1)
	target.Lsh(target, uint(256-bits))
	fmt.Println(target.String())
	nonce := 0
	serviceStr := "block data"
	var hashInt big.Int
	for {
		data := serviceStr + strconv.Itoa(nonce)
		hash := sha256.Sum256([]byte(data))
		hashInt.SetBytes(hash[:])
		fmt.Println(hashInt.String(), nonce)
		if hashInt.Cmp(target) == -1 {
			fmt.Println("本机挖矿成功")
			return
		}
		nonce++
	}

}
