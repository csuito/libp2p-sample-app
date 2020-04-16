package core

import (
	"time"

	"github.com/csuito/block/core/types"
)

var (
	chain *Blockchain = &Blockchain{}
)

type Blockchain []types.Block

func Get() *Blockchain {
	return chain
}

func (bc *Blockchain) Init() error {
	gb := types.Block{}
	gb = types.Block{
		Index:     0,
		Timestamp: time.Now().UTC().String(),
		Message:   "Genesis",
		Hash:      gb.CalculateHash(),
		PrevHash:  "",
	}
	if err := chain.AddBlock(gb); err != nil {
		return err
	}
	return nil
}

func (bc *Blockchain) AddBlock(b types.Block) error {
	if len(*bc) > 0 {
		if err := b.Validate((*bc)[len(*bc)-1]); err != nil {
			return err
		}
	}
	*bc = append(*bc, b)
	return nil
}
