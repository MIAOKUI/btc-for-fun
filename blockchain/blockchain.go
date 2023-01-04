package blockchain

type BlockChain struct {
	lashHash Hash
	blocks   map[Hash]*Block
}

func NewBlockChain() *BlockChain {
	bc := &BlockChain{
		blocks: map[Hash]*Block{},
	}
	return bc
}

func (bc *BlockChain) AddGenesisBlock() *BlockChain {
	if bc.lashHash != "" {
		return bc
	}
	return bc.AddBlock("The Genesis Block")
}

func (bc *BlockChain) AddBlock(txs string) *BlockChain {
	b := NewBlock(bc.lashHash, txs)
	bc.blocks[b.hashCurr] = b
	bc.lashHash = b.hashCurr
	return bc
}
