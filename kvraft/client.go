package kvraft

import (
	"crypto/rand"
	"math/big"

	"../labrpc"
)

// Clerk defines the type for managing
// RPC interactions with the servers.
type Clerk struct {
	servers []*labrpc.ClientEnd
	// You will have to modify this struct.
}

// nrand returns a random number
func nrand() int64 {
	max := big.NewInt(int64(1) << 62)
	bigx, _ := rand.Int(rand.Reader, max)
	x := bigx.Int64()
	return x
}

// MakeClerk creates the clerk agent
// for RPC interactions with the servers.
func MakeClerk(servers []*labrpc.ClientEnd) *Clerk {
	ck := new(Clerk)
	ck.servers = servers
	// You'll have to add code here.
	return ck
}

// Get fetches the current value for a key.
// returns "" if the key does not exist.
// keeps trying forever in the face of all other errors.
//
// you can send an RPC with code like this:
// ok := ck.servers[i].Call("KVServer.Get", &args, &reply)
//
// the types of args and reply (including whether they are pointers)
// must match the declared types of the RPC handler function's
// arguments. and reply must be passed as a pointer.
//
func (ck *Clerk) Get(key string) string {
	value := ""
	nServers := int64(len(ck.servers))
	server := nrand() % nServers // Randomly choose a server
	args := &GetArgs{
		Key: key,
	}

	for {
		// Send Get RPC
		reply := &GetReply{}
		if ok := ck.servers[server].Call("KVServer.Get", args, reply); !ok {
			// If cannot reach server continue trying
			continue
		}

		// No such key, return empty string
		if reply.Err == ErrNoKey {
			break
		} else if reply.Err == ErrWrongLeader {
			server = nrand() % nServers // Randomly choose new server (choose leader based on reply?)
		} else {
			value = reply.Value
			break
		}

	}
	return value
}

// PutAppend method performs put or append.
//
// you can send an RPC with code like this:
// ok := ck.servers[i].Call("KVServer.PutAppend", &args, &reply)
//
// the types of args and reply (including whether they are pointers)
// must match the declared types of the RPC handler function's
// arguments. and reply must be passed as a pointer.
//
func (ck *Clerk) PutAppend(key string, value string, op string) {
	// You will have to modify this function.
}

// Put set or update a key's value.
func (ck *Clerk) Put(key string, value string) {
	ck.PutAppend(key, value, "Put")
}

// Append adds arg to key's value.
// Append to a non-existant key acts like Put.
func (ck *Clerk) Append(key string, value string) {
	ck.PutAppend(key, value, "Append")
}
