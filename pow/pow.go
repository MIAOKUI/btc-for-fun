package pow

import (
	"crypto/sha256"
	"github.com/MIAOKUI/btc-for-fun/blockchain"
	"math"
	"math/big"
	"strconv"
)

type ProofOfWork struct {
	block  *blockchain.Block
	target *big.Int
}

func NewPow(b *blockchain.Block) *ProofOfWork {
	p := &ProofOfWork{
		block:  b,
		target: nil,
	}
	target := big.NewInt(1)
	target.Lsh(target, uint(blockchain.HashLen-b.GetBits()+1))
	p.target = target
	return p
}

func Proof(p *ProofOfWork) (int, string) {
	var hashInt big.Int
	serviceString := p.block.GetServiceString()
	nonce := 0
	for nonce < math.MaxInt64 {
		hash := sha256.Sum256([]byte(serviceString + strconv.Itoa(nonce)))
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(p.target) == -1 {
			return nonce, serviceString
		}
		nonce++
	}
	return 0, ""
}
