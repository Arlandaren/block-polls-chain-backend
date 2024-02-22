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

func calculate_hash(index uint64, prev string, timestamp int,hash string,data string,owner uint64) string{
	value := strconv.FormatUint(index,10)+ prev +strconv.FormatInt(int64(timestamp),10) + data +strconv.FormatUint(owner,10)
	hash_value := sha256.New()
	hash_value.Write([]byte(value))
	hashed := hash_value.Sum(nil)
	return hex.EncodeToString(hashed)
}

func AddBlock(data string, owner uint64) *Block{
	lastblock := GetLastBlock()
	return lastblock
}
