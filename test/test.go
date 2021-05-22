package main

//
// start the coordinator process, which is implemented
// in ../mr/coordinator.go
//
// go run mrcoordinator.go pg*.txt
//
// Please do not change this file.
//

import "os"
import "encoding/json"
import "fmt"
import "sync"

func main() {
//        hey , _ := os.Create("/home/se/distrib/test/hey.json")
 //       enc := json.NewEncoder(hey)
  //      enc.Encode(map[string]int{"apple":5,"lettuce":7})
    var wg sync.WaitGroup
    wg.Add(1)
    go hey(&wg)
    wg.Add(1)
    go hey(&wg)
    wg.Wait()


}

func hey(wg *sync.WaitGroup){
        defer wg.Done()
        hey2 , _ := os.OpenFile("hey.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
        enc2 := json.NewEncoder(hey2)
        enc2.Encode(map[string]int{"apple":5,"lettuce":7})
        fmt.Println("hey")

}
