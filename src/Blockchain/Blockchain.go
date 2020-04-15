package Blockchain

import (
	. "Block"
	. "Miner"
	//"fmt"
	"math"
	"strconv"
)

type Blockchain struct {
	numBlocks 	int
	genesis		*Block
	topBlock	*Block 		// the newest block..
	miners 		[]Miner
}

func (bc *Blockchain) SerialRun(start_hash string, num_blocks int) {
	for i:=0; i<num_blocks; i++ {
		buffer_channel := make(chan Block)

		// distribute a new empty block to all the miners..
		new_block := InitBlock(bc.GetNumBlocks(), bc.topBlock.GetHash())
		for i:=0; i<len(bc.miners); i++ {
			go bc.miners[i].SerialMine(InitBlock(bc.GetNumBlocks(), bc.topBlock.GetHash()), start_hash, buffer_channel)
		}

		// Get the first response because that's all that matters...
		select {
			case new_block = <- buffer_channel:
				bc.AddBlock(&new_block)
		}
	}
}


func (bc *Blockchain) ParallelRun(start_hash string, num_blocks int) {
	for j:=0; j<num_blocks; j++ {
		buffer_channel := make(chan Block)

		// Get the total number of devices that can perform calculations..
		num_devices := 0
		for i:=0; i<len(bc.miners); i++ {
			num_devices = num_devices + bc.miners[i].GetNumDevices()
		}

		// calculate the range that each device should operate over
		range_per_device := math.MaxInt32/num_devices

		// create the new block for later...
		new_block := InitBlock(bc.GetNumBlocks(), bc.topBlock.GetHash())

		// distribute the starting and ending range as well as a new empty block to all the miners..
		start_range := 0
		for i:=1; i<=len(bc.miners); i++ {
			end_range := start_range + (bc.miners[i-1].GetNumDevices()*range_per_device) - 1
			go bc.miners[i-1].ParallelMine(InitBlock(bc.GetNumBlocks(), bc.topBlock.GetHash()), start_hash, start_range, end_range, buffer_channel)

			start_range = end_range
		}

		// Get the first response because that's all that matters...
		select {
			case new_block = <- buffer_channel:
				bc.AddBlock(&new_block)
		}
	}
}



func (bc *Blockchain) AddBlock(b *Block) {
	// link the blockchain's topBlock to the newest block created.
	bc.topBlock.SetNextBlock(b)
	bc.topBlock = b

	bc.IncrementNumBlocks()
}

func (bc *Blockchain) AddMiners(m []Miner) {
	bc.miners = m
}

func InitBlockchain() *Blockchain{
	var miners []Miner
	// number blocks is 1 because we create the genesis block.
	genesis := InitBlock(0, "GENESIS")
	return &Blockchain{1, &genesis, &genesis, miners}
}

func (bc *Blockchain) GetNumBlocks() int {
	return bc.numBlocks
}

func (bc *Blockchain) GetGenesis() *Block {
	return bc.genesis
}

func (bc *Blockchain) GetTopBlock() *Block {
	return bc.topBlock
}

// this could have an off by 1 error. 
func (bc *Blockchain) GetBlock(index int) *Block{
	curr_block := bc.genesis
	for i:=0; i<index; i++ {
		curr_block = curr_block.GetNextBlock()
	}

	return curr_block
}

func (bc *Blockchain) GetNumMiners() int {
	return len(bc.miners)
}

func (bc *Blockchain) IncrementNumBlocks() {
	bc.numBlocks++
}


func (bc *Blockchain) ToString() string {
	s :="Number Miners = " + strconv.Itoa(bc.GetNumMiners()) + "\n"
	s = s + "Num Blocks = " + strconv.Itoa(bc.GetNumBlocks()) + "\n\n"

	for i:=0; i<bc.GetNumBlocks(); i++ {
		curr_block := bc.GetBlock(i)
		s = s + curr_block.ToString() + "\n\n"
	}

	return s
}