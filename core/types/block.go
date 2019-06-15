package types

import (
	"ADS/core/merkleTree"
	"bytes"
	"crypto/sha256"
	"encoding"

	"fmt"

	"crypto"
	"time"
)

type Header struct {
	ParentHash 	string
	Version    int8
	MerkleRoot crypto.Hash
	Timestamp	time.Time

}

type Block  struct {
	ParentHash 	string
	MsgHash []message
}

func GeneisBlock() *Block{
	return &Block{

	}
}


//计算默克尔树
func (b Block) MerkleRoot() crypto.Hash {





}


	//添加新区块
func NewBlock(header *Header,msgs []*message) *Block{



	return &Block{

	}
}

func (b Block) PrevHash()string{
	return b.ParentHash
}

//计算当前区块Hash
func (bh Header) CalculateHashForBlock() string{
	return fmt.Sprintf("%x",sha256.Sum256([]byte(fmt.Sprintf("s%d%s%d%",bh.ParentHash,bh.MerkleRoot,bh.Timestamp,bh.Version))))
}

