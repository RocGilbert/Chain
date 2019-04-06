package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64 //区块编号
	Timestamp int64  //区块时间戳
	PrevBlockHash string //上一个区块的哈希值
	Hash string //当前区块哈希值

	Data string //区块数据

}

//定义一个函数 来计算哈希值   以字符串来表示的哈希值
func calculateHash(b Block) string {
	//blockData字符串类型
	blockData := string(b.Index) + string(b.Timestamp)+ b.PrevBlockHash + b.Data
	//Sum256 将字节切片转换为字节数组  返回的是字节数组
	//[]byte(blockData) 将blockData字符串类型  转换为 32个元素的字节切片
	//hashInBytes 字节数组
	 hashInBytes := sha256.Sum256([]byte(blockData))

	 //EncodeToString 要求里面是字节切片   hashInBytes[:] 将数组 转换为字节切片
	 return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
	//定义一个新区块
	newBlock := Block{}
	//新区块
	newBlock.Index = preBlock.Index + 1
	newBlock.Data = data
	//新区块的上一个区块的哈希值 等于 上一个区块的哈希值
	newBlock.PrevBlockHash = preBlock.Hash
	//新区块的时间 等于 当前时间
	newBlock.Timestamp = time.Now().Unix()
	//新区块的哈希值 是 计算得来的
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

//用来生成创世区块。
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock,"Genisis Block")
}