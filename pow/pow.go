package pow

import (
	"crypto/sha256"
	"fmt"
	"github.com/MIAOKUI/btc-for-fun/block"
	"math"
	"math/big"
	"strconv"
)

type ProofOfWork struct {
	block  *block.Block
	target *big.Int
}

func NewPow(b *block.Block) *ProofOfWork {
	p := &ProofOfWork{
		block: b,
		//target: nil,
	}
	target := big.NewInt(1)
	target.Lsh(target, uint(block.HashLen-b.GetBits()+1))
	p.target = target
	return p
}

func (p *ProofOfWork) Proof() (int, block.Hash) {
	var hashInt big.Int
	serviceString := p.block.GetServiceString()
	nonce := 0
	fmt.Printf("Target:%d\n", p.target)
	for nonce < math.MaxInt64 {
		hash := sha256.Sum256([]byte(serviceString + strconv.Itoa(nonce)))
		hashInt.SetBytes(hash[:])
		fmt.Printf("hash: %s\n", hashInt.String())
		if hashInt.Cmp(p.target) == -1 {
			return nonce, fmt.Sprintf("%x", hash)
		}
		nonce++
	}
	return 0, ""
}
