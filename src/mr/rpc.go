package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}



// Add your RPC definitions here.

type TaskReqArgs struct {
    // the coordinator doesn't need to know much abt the worker so prolly not much here
}

type TaskReqReply struct {
	Err Err
    assignment, filename, map_part, red_hash string
    nReduce int
    // might want to make file and key the same
    // file and contents are for map
    // key and values are for reduce
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
