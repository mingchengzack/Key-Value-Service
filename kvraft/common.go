package kvraft

import "log"

const (
	// OK means valid
	OK = "OK"

	// ErrNoKey defines the error for non-existant key
	ErrNoKey = "ErrNoKey"

	// ErrWrongLeader defines the error for contacting wrong server
	ErrWrongLeader = "ErrWrongLeader"
)

// Debug mode
const Debug = 0

// Err structure is just string
type Err string

// PutAppendArgs is the structure for handling put or append RPC
// from client
type PutAppendArgs struct {
	Key   string
	Value string
	Op    string // "Put" or "Append"
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
}

// PutAppendReply is the structure for storing put or append RPC's
// reply from server
type PutAppendReply struct {
	Err Err
}

// GetArgs is the structure for handling get RPC
// from client
type GetArgs struct {
	Key string
	// You'll have to add definitions here.
}

// GetReply is the structure for storing get RPC's
// reply from server
type GetReply struct {
	Err   Err
	Value string
}

// DPrintf for debug printf
func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug > 0 {
		log.Printf(format, a...)
	}
	return
}
