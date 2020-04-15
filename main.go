package main

import (
	. "Blockchain"
	. "Miner"
    "fmt"
    "strconv"
)

func main() {
	max_miners := 5
	blockchain := InitBlockchain()

	miners := make([]Miner, max_miners)

	for i:=0; i<max_miners; i++ {
		miners[i] = CreateMiner("Miner" + strconv.Itoa(i), 5)
	}

	blockchain.AddMiners(miners)

	blockchain.SerialRun("00000", 5)
	//blockchain.ParallelRun("00000", 5)
	
	fmt.Println(blockchain.ToString())
}