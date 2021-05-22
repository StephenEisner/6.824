package mr

import "log"
import "net"
import "os"
import "net/rpc"
import "net/http"


type Coordinator struct {
	// Your definitions here.

    nreduce int
    //bool has if that file has been fully processed
    file_map map[string]bool
    //tasks still in progress
    in_progress []interface{}

}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}


func (c *Coordinator) Assign(args *TaskReqAsk, reply *TaskReqReply) error {
    // TODO fill reply with the relevant information
    // TODO add args to a list of tasks currently outstanding
    // rremember to lock the reply when you are modifying it
	return nil
    //change to return an error if there are no tasks available
}



//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
    // ret := false
    // return ret

	// Your code here.

    //check if list of completed files has everything completed
    for _,done:= range c.file_map {
        if !done{
            return false
        }
    }

	return true
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.

    //partition input files

    //set nReduce
    c.nReduce = nReduce

    //create list of completed files
    for _, f := range files{
        c.file_map[f] = false
    }

    //create empty list of current workers

	c.server()
	return &c
}
