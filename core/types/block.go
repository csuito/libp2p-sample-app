package types

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Message   string
	Hash      string
	PrevHash  string
}

func (b *Block) Validate(lastBlock Block) error {
	if lastBlock.Index+1 != b.Index {
		return errors.New("invalid index")
	}
	if lastBlock.Hash != b.PrevHash {
		return errors.New("invalid prevHash")
	}
	if b.Hash != b.CalculateHash() {
		return errors.New("invalid hash")
	}
	return nil
}

func (b *Block) CalculateHash() string {
	r := strconv.Itoa(b.Index) + b.Timestamp + b.Message + b.PrevHash
	h := sha256.New()
	h.Write([]byte(r))
	return hex.EncodeToString(h.Sum(nil))
}

func NewBlock(prevBlock Block, msg string) Block {
	b := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().UTC().String(),
		Message:   msg,
		PrevHash:  prevBlock.Hash,
	}
	b.Hash = b.CalculateHash()
	return b
}
