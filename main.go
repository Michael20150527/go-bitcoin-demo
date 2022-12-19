package main

import "fmt"

//创世语
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

	return &block
}

func main() {
	fmt.Println("hello world")
	block := NewBlock(genesisInfo, []byte{0x0000000000000000})

	fmt.Printf("PrevBlockHash : %x\n", block.PrevBlockHash)
	fmt.Printf("Hash : %x\n", block.Hash)
	fmt.Printf("Data : %s\n", block.Data)
}
