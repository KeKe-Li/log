package v1

import (
	"sync"

	"github.com/KeKe-Li/log/uuid/rand"
)

//   +------ 0 ------+------ 1 ------+------ 2 ------+------ 3 ------+
//   +0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1+
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                          time_low                             |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |       time_mid                |         time_hi_and_version   |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |clk_seq_hi_res |  clk_seq_low  |         node (0-1)            |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                         node (2-5)                            |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

var node = rand.MAC[:] // read only.

const sequenceMask uint32 = 0x3FFF // 14bits

var (
	gMutex         sync.Mutex
	gSequenceStart uint32 = rand.Uint32() & sequenceMask
	gLastTimestamp int64  = -1
	gLastSequence  uint32 = gSequenceStart
)

// New returns a STANDARD version 1 UUID.
func New() (uuid [16]byte) {
	var (
		timestamp = uuidTimestamp()
		sequence  uint32
	)

	gMutex.Lock() // Lock
	switch {
	case timestamp > gLastTimestamp:
		sequence = gSequenceStart
		gLastTimestamp = timestamp
		gLastSequence = sequence
		gMutex.Unlock() // Unlock
	case timestamp == gLastTimestamp:
		sequence = (gLastSequence + 1) & sequenceMask
		if sequence == gSequenceStart {
			timestamp = tillNext100nano(timestamp)
			gLastTimestamp = timestamp
		}
		gLastSequence = sequence
		gMutex.Unlock() // Unlock
	default: // timestamp < lastTimestamp
		gSequenceStart = rand.Uint32() & sequenceMask // NOTE
		sequence = gSequenceStart
		gLastTimestamp = timestamp
		gLastSequence = sequence
		gMutex.Unlock() // Unlock
	}

	// time_low
	uuid[0] = byte(timestamp >> 24)
	uuid[1] = byte(timestamp >> 16)
	uuid[2] = byte(timestamp >> 8)
	uuid[3] = byte(timestamp)

	// time_mid
	uuid[4] = byte(timestamp >> 40)
	uuid[5] = byte(timestamp >> 32)

	// time_hi_and_version
	uuid[6] = byte(timestamp>>56) & 0x0F
	uuid[6] |= 0x10 // version 1, 4bits
	uuid[7] = byte(timestamp >> 48)

	// clk_seq_hi_res
	uuid[8] = byte(sequence>>8) & 0x3F
	uuid[8] |= 0x80 // variant, 2bits

	// clk_seq_low
	uuid[9] = byte(sequence)

	// node
	copy(uuid[10:], node)
	return
}
