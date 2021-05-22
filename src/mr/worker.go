package mr

import "fmt"
import "log"
import "net/rpc"
import "hash/fnv"
import "time"


//
// Map functions return a slice of KeyValue.
//
type KeyValue struct {
	Key   string
	Value string
}

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}


//
// main/mrworker.go calls this function.
//
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {
        for {
            // Your worker implementation here.
            args := TaskReqArgs{}
            reply := TaskReqReply{}

            call("Coordinator.Assign", &args, &reply)

            //Execute what is given in reply
            switch reply.assignment (
            case 'map':
                    worker_map(mapf, &reply)
                    //complete with behavior when a worker is tasked with map
            case 'reduce':
                    worker_reduce(reducef, &reply)
                    //complete with behavior when a worker is tasked with reduce
            case 'wait':
                    time.Sleep(5 * time.Second)
            case 'quit':
                    return
            default:
                    panic("invalid worker  assignment")
            )

        }
}

//THINGS TO CONSIDER.
//might need to look the rpcs when they are being read or modified. maybe it is fine since they are def not going to be double read but is that bad form?

func worker_map(mapf func(string, string) []KeyValue, reply *TaskReqReply) {
    file, err := os.Open(reply.filename)
    if err != nil {
            log.Fatalf("cannot open %v", filename)
    }
    content, err := ioutil.ReadAll(file)
    if err != nil {
            log.Fatalf("cannot read %v", filename)
    }
    file.Close()
    kva := mapf(reply.filename, string(content))
    for kv := range kva {

            //reduce group
            red_grp := ihash(kv.Key) % reply.nReduce

            filename := "mr-" + strconv.Itoa(reply.map_part) + "-" + strconv.Itoa(red_grp)

            f , err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
            if err != nil {
                        panic(err)
                }
            defer f.Close()

            enc := json.NewEncoder(f)
            err := enc.Encode(&kv)
    }

}

//TODO current contents are just the sequential reduce section this is definitely not correct
func worker_reduce(mapf reducef func(string, []string) string, reply *TaskReqReply) {
	i := 0
	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}
		values := []string{}
		for k := i; k < j; k++ {
			values = append(values, intermediate[k].Value)
		}
		output := reducef(intermediate[i].Key, values)

		// this is the correct format for each line of Reduce output.
		fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)

		i = j
	}



}
//
// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
