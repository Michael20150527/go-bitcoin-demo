package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func uintToByte(num uint64) []byte {
	// 使用binary.Write进行编码
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}
