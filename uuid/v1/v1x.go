package v1

import (
	"os"
	"sync"

	"github.com/KeKe-Li/log/uuid/rand"
)

//   +------ 0 ------+------ 1 ------+------ 2 ------+------ 3 ------+
//   +0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1+
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                          time_low                             |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |       time_mid                |        time_hi_and_pid_low    |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |clk_seq_hi_pid |  clk_seq_low  |         node (0-1)            |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                         node (2-5)                            |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

var pid = byte(hash(uint64(os.Getpid()))) // 6-bit hash of os.Getpid(), read only.

// hash uint64 to a 6-bit integer value.
func hash(x uint64) uint64 {
	return (x ^ x>>6 ^ x>>12 ^ x>>18 ^ x>>24 ^ x>>30 ^ x>>36 ^ x>>42 ^ x>>48 ^ x>>54 ^ x>>60) & 0x3f
}

var xNode = MAC[:] // read only.

const xSequenceMask uint32 = 0x3FFF // 14bits

var (
	gxMutex         sync.Mutex
	gxSequenceStart uint32 = rand.Uint32() & xSequenceMask
	gxLastTimestamp int64  = -1
	gxLastSequence  uint32 = gxSequenceStart
)

// Newx returns a NONSTANDARD UUID(lower probability of conflict).
func Newx() (uuid [16]byte) {
	var (
		timestamp = uuidTimestamp()
		sequence  uint32
	)

	gxMutex.Lock() // Lock
	switch {
	case timestamp > gxLastTimestamp:
		sequence = gxSequenceStart
		gxLastTimestamp = timestamp
		gxLastSequence = sequence
		gxMutex.Unlock() // Unlock
	case timestamp == gxLastTimestamp:
		sequence = (gxLastSequence + 1) & xSequenceMask
		if sequence == gxSequenceStart {
			timestamp = tillNext100nano(timestamp)
			gxLastTimestamp = timestamp
		}
		gxLastSequence = sequence
		gxMutex.Unlock() // Unlock
	default: // timestamp < xLastTimestamp
		gxSequenceStart = rand.Uint32() & xSequenceMask // NOTE
		sequence = gxSequenceStart
		gxLastTimestamp = timestamp
		gxLastSequence = sequence
		gxMutex.Unlock() // Unlock
	}

	// time_low
	uuid[0] = byte(timestamp >> 24)
	uuid[1] = byte(timestamp >> 16)
	uuid[2] = byte(timestamp >> 8)
	uuid[3] = byte(timestamp)

	// time_mid
	uuid[4] = byte(timestamp >> 40)
	uuid[5] = byte(timestamp >> 32)

	// time_hi_and_pid_low
	uuid[6] = byte(timestamp >> 52)
	uuid[7] = byte(timestamp>>48) << 4
	uuid[7] |= pid & 0x0F // pid, 4bits

	// clk_seq_hi_pid
	uuid[8] = byte(sequence>>8) & 0x3F
	uuid[8] |= (pid << 2) & 0xC0 // // pid, 2bits

	// clk_seq_low
	uuid[9] = byte(sequence)

	// node
	copy(uuid[10:], xNode)
	return
}
