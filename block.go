package main

import "crypto/sha256"

// 创世语
const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

type Block struct {
	PrevBlockHash []byte // 前区块哈希
	Hash          []byte // 当前区块哈希
	Data          []byte // 数据，目前使用字节流，V4开始使用交易代替
}

// 创建区块，对Block的每一个字段填充数据即可
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{}, // 先填充为空，后续会填充数据
		Data:          []byte(data),
	}

	block.setHash()

	return &block
}

// 为了生成区块哈希，我们实现一个简单函数，来计算哈希值，没有随机数，没有难度值
func (block *Block) setHash() {
	var data []byte
	data = append(data, block.PrevBlockHash...)
	data = append(data, block.Data...)

	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

// 创建区块链，使用Block数组模拟
type BlockChain struct {
	Blocks []*Block
}
