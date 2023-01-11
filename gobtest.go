package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Block struct {
	CurrHash string
	TXs      string
}

func main() {
	b := Block{
		CurrHash: "fdsafdsafdsafdsa",
		TXs:      "first transaction",
	}
	var bb bytes.Buffer
	enc := gob.NewEncoder(&bb)

	err := enc.Encode(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bb.Bytes(), bb.String())

	var bbr bytes.Buffer
	bbr.Write(bb.Bytes())
	b1 := Block{}
	dec := gob.NewDecoder(&bbr)
	dec.Decode(&b1)
	fmt.Println(dec)
}
