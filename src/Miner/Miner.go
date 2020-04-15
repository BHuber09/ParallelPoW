package Miner

import ( 
	. "Block"
	"crypto/sha1"
	"encoding/base64"
	"math"
)

type Miner struct {
	minerId 	string
	numDevices	int
}


func isChannelClosed(ch <- chan Block) bool {
	select {
	case <- ch:
		return true
	default: 
	}

	return false
}

// Leading hash is going to be the 'puzzle' part of the POW..
func (m Miner) SerialMine(b Block, leading_hash string, buffer_channel chan Block) {
	b.Miner = m.GetMinerId()

	// hash the block.
	hasher := sha1.New()

	for i := 0; i < math.MaxInt32; i++ {
		// first we test to see if the channel is closed...
		// if(isChannelClosed(buffer_channel)) {
		// 	return
		// }

		// Change the header in hopes to get a hit on the puzzle. 
		b.Header = i

		// calculate and compare the hashes.
		hasher.Write([]byte(b.ToString()))
		hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		// substringing to get the same sized strings
		sub_hash := hash[0:len(leading_hash)]

		if(leading_hash == sub_hash) {
			// found it... Return the block.
			b.Hash = hash
			
			// check if it's closed again right before we push data to it..
			//if(isChannelClosed(buffer_channel)) {
			buffer_channel <- b
			//}
			
			break
		}
	}
}

func (m Miner)ParallelMine(b Block, leading_hash string, start_range int, end_range int, buffer_channel chan Block) {
	b.Miner = m.GetMinerId()

	// hash the block.
	hasher := sha1.New()

	for i := start_range; i < end_range; i++ {
		// first we test to see if the channel is closed...
		// if(isChannelClosed(buffer_channel)) {
		// 	return
		// }

		// Change the header in hopes to get a hit on the puzzle. 
		b.Header = i

		// calculate and compare the hashes.
		hasher.Write([]byte(b.ToString()))
		hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		// substringing to get the same sized strings
		sub_hash := hash[0:len(leading_hash)]

		if(leading_hash == sub_hash) {
			// found it... Return the block.
			b.Hash = hash

			// check if it's closed again right before we push data to it..
			//if(isChannelClosed(buffer_channel)) {
			buffer_channel <- b
			//}

			break
		}
	}
}	

func CreateMiner(_minerId string, _numDevices int) Miner {
	return Miner{_minerId, _numDevices}
}

func (m Miner) GetMinerId() string {
	return m.minerId
}

func (m Miner) GetNumDevices() int {
	return m.numDevices
}