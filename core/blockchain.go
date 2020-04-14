package core

import (
	"time"

	"github.com/csuito/block/core/types"
)

type blockchain []types.Block

func Initialize() *blockchain {
	chain := &blockchain{}
	gb := types.Block{}
	gb = types.Block{
		Index:     0,
		Timestamp: time.Now().UTC().String(),
		BPM:       0,
		Hash:      gb.CalculateHash(),
		PrevHash:  "",
	}
	*chain = append(*chain, gb)
	return chain
}

func (bc *blockchain) AddBlock(b types.Block) error {
	if len(*bc) > 0 {
		if err := b.Validate((*bc)[len(*bc)-1]); err != nil {
			return err
		}
	}
	*bc = append(*bc, b)
	return nil
}
