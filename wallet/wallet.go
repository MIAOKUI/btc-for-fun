package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type Address = string

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    *Address
}

func NewWallet() *Wallet {
	w := &Wallet{}
	w.GenKey()
	w.GenAddress()
	return w
}

func (w *Wallet) GenKey() *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	w.privateKey = privateKey
	return w
}

func (w *Wallet) GenAddress() *Wallet {
	pubKey := w.genPubKey()

}

func (w *Wallet) genPubKey() []byte {
	pubKey := append(w.privateKey.PublicKey.X.Bytes(), w.privateKey.PublicKey.Y.Bytes()...)
	return pubKey
}
