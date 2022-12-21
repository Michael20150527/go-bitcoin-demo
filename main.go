package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// block := NewBlock(genesisInfo, []byte{0x0000000000000000})

	bc := NewBlockChain()
	bc.AddBlock("班主任来了，大家欢迎~")
	bc.AddBlock("班主任走了~")

	for i, block := range bc.Blocks {
		fmt.Printf("+++++++++++++++%d++++++++++++++++\n", i)
		fmt.Printf("PrevBlockHash : %x\n", block.PrevBlockHash)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Data : %s\n", block.Data)
	}
}
