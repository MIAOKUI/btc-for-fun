package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

const HashLen = 256

type BlockData struct {
	Version        int
	HashPrevBlock  Hash
	HashMerkleRoot Hash
	Time           time.Time
	Bits           int
	Nonce          int
	Txs            string
	TxCounter      int
	HashCurr       Hash
}

func BlockSerialize(b Block) ([]byte, error) {
	bd := BlockData{
		Version:        b.header.version,
		HashPrevBlock:  b.header.hashPrevBlock,
		HashMerkleRoot: b.header.hashMerkleRoot,
		Time:           b.header.time,
		Bits:           b.header.bits,
		Nonce:          b.header.nonce,
		Txs:            b.txs,
		TxCounter:      b.txCounter,
		HashCurr:       b.hashCurr,
	}
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(bd); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func BlockDeserialize(bsData []byte) (Block, error) {
	buffer := bytes.Buffer{}
	buffer.Write(bsData)
	decoder := gob.NewDecoder(&buffer)
	bd := &BlockData{}

	if err := decoder.Decode(bd); err != nil {
		return Block{}, err
	}
	return Block{
		header: BlockHeader{
			version:        bd.Version,
			hashPrevBlock:  bd.HashPrevBlock,
			hashMerkleRoot: bd.HashMerkleRoot,
			time:           bd.Time,
			bits:           bd.Bits,
			nonce:          bd.Nonce,
		},
		txs:       bd.Txs,
		txCounter: bd.TxCounter,
		hashCurr:  bd.HashCurr,
	}, nil
}
