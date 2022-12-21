package main

import (
	"crypto/sha256"
	"time"
)

// 创世语
const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

type Block struct {
	Version       uint64 // 区块版本号
	PrevBlockHash []byte // 前区块哈希
	MerkleRoot    []byte // 先填写为空，后续v4的时候使用
	TimeStamp     uint64 // 从1970.1.1至今的秒数
	Difficulty    uint64 // 挖矿的难度值，v2时使用
	Nonce         uint64 // 随机数，挖矿找的就是它
	Hash          []byte // 当前区块哈希
	Data          []byte // 数据，目前使用字节流，V4开始使用交易代替
}

// 创建区块，对Block的每一个字段填充数据即可
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:       00,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    []byte{},
		TimeStamp:     uint64(time.Now().Unix()),
		Difficulty:    10,
		Nonce:         10,
		Hash:          []byte{}, // 先填充为空，后续会填充数据
		Data:          []byte(data),
	}

	block.setHash()

	return &block
}

// 为了生成区块哈希，我们实现一个简单函数，来计算哈希值，没有随机数，没有难度值
func (block *Block) setHash() {
	var data []byte
	data = append(data, uintToByte(block.Version)...)
	data = append(data, block.PrevBlockHash...)
	data = append(data, block.MerkleRoot...)
	data = append(data, uintToByte(block.TimeStamp)...)
	data = append(data, uintToByte(block.Difficulty)...)
	data = append(data, uintToByte(block.Nonce)...)
	data = append(data, block.Data...)

	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

// 创建区块链，使用Block数组模拟
type BlockChain struct {
	Blocks []*Block
}
