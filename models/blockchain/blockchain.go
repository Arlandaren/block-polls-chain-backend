package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct{
	Index uint32
	Prev_hash string
	Timestamp time.Time
	Hash string
	Data string
	Owner uint64
}

func calculate_hash(index uint64, prev string, timestamp time.Time,data string,owner uint64) string{
	value := strconv.FormatUint(index,10)+ prev +strconv.FormatInt((timestamp.UnixNano()),10) + data +strconv.FormatUint(owner,10)
	hash_value := sha256.New()
	hash_value.Write([]byte(value))
	hashed := hash_value.Sum(nil)
	return hex.EncodeToString(hashed)
}

func AddBlock(data string, owner uint64) *Block{
	lastblock := GetLastBlock()
	var block Block
	block.Index = lastblock.Index + 1
	block.Prev_hash = lastblock.Hash
	block.Timestamp = time.Now()
	block.Hash = calculate_hash(uint64(block.Index),lastblock.Hash,block.Timestamp,data,owner)
	block.Data = data
	block.Owner = owner

	return &block
}
