package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct{
	Index uint32
	PreviousHash string
	Timestamp time.Time
	Hash string
	Data string
}

func calculate_hash(index uint64, prev string, timestamp time.Time,data string) string{
	value := strconv.FormatUint(index,10)+ prev +strconv.FormatInt((timestamp.UnixNano()),10) + data
	hash_value := sha256.New()
	hash_value.Write([]byte(value))
	hashed := hash_value.Sum(nil)
	return hex.EncodeToString(hashed)
}

func AddBlock(data string) (*Block, error){
	lastblock := GetLastBlock()
	var block Block
	block.Index = lastblock.Index + 1
	block.PreviousHash = lastblock.Hash
	block.Timestamp = time.Now()
	block.Hash = calculate_hash(uint64(block.Index),lastblock.Hash,block.Timestamp,data)
	block.Data = data
	err := InsertBlockIntoDb(&block)
	if err != nil{
		return nil,err
	}
	return &block, nil
}
