package Block

import ( 
	"strconv"
	"time"
)

/*
This is the individualized block
	Index is the current index that this block is in the blockckain
	Timestamp is the current time when the block is being made
	Data is the transactional data of the block
	Hash is this block's hash
	PrevHash is the hash of the previous block
	Header is going to be the data that is modified to minimize the hash

*/
type Block struct {
	Index		int
	Timestamp	string
	Data 		int
	Hash 		string
	PrevHash 	string
	Header		int
	Miner 		string
	nextBlock	*Block
}

func InitBlock(index int, prev_hash string) Block {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	return Block{index, timestamp, index, "nil", prev_hash, 0, "nil", nil}
}

func (b *Block) GetIndex() int {
	return b.Index
} 

func (b *Block) GetTimestamp() string{
	return b.Timestamp
} 

func (b *Block) GetData() int{
	return b.Data
} 

func (b *Block) GetHash() string{
	return b.Hash
} 

func (b *Block) GetPrevHash() string{
	return b.PrevHash
} 

func (b *Block) GetHeader() int{
	return b.Header
}

func (b *Block) GetMiner() string{
	return b.Miner
}

func (b *Block) GetNextBlock() *Block{
	return b.nextBlock
}

func (b *Block) SetNextBlock(next_b *Block) {
	b.nextBlock = next_b
}

func (b *Block) ToString() string{
	s := "Block: " + strconv.Itoa(b.GetIndex()) + "\n"
	s = s + "Timestamp: " + b.GetTimestamp() + "\n"
	s = s + "Data: " + strconv.Itoa(b.GetData()) + "\n"
	s = s + "Hash: " + b.GetHash() + "\n"
	s = s + "PrevHash: " + b.GetPrevHash() + "\n"
	s = s + "Header: " + strconv.Itoa(b.GetHeader()) + "\n"
	s = s + "Miner: " + b.GetMiner()

	return s
}