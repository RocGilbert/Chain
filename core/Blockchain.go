package core

import "fmt"

//定义一个Blockchain结构体
type Blockchain struct {
	//Block 数组切片
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}

func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := GenerateNewBlock(*preBlock,data)
	bc.AppendBlock(&newBlock)
}

//Blockchain 结构体的方法  传进去一个
func (bc *Blockchain) AppendBlock(newBlock *Block)  {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if(isValid(*newBlock, *bc.Blocks[len(bc.Blocks) - 1])){
		bc.Blocks = append(bc.Blocks,newBlock)
	}

}

func (bc *Blockchain) Print() {
	for _,block := range bc.Blocks{
		fmt.Printf("Index:%d\n",block.Index)
		fmt.Printf("Prev.Hash:%s\n",block.PrevBlockHash)
		fmt.Printf("Curr.Hash:%s\n",block.Hash)
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("Timestamp:%d\n",block.Timestamp)
		fmt.Println()
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index{
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash{
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}