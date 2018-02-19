package main

import (
	"bufio"
	//"log"
	"os"
	"fmt"
	"runtime"
	//"time"
	//"syscall"
	"strings"
)

/*func show_thread_id (a int) {
	fmt.Printf("syscall.Gettid routine%d: %v\n",a,syscall.Gettid())
}*/

//curl https://golang.org | xargs | go run main.go
func main() {

	var numCPU = runtime.NumCPU()

	//lock if it is more than N senders or more than N receivers
	var block chan int = make(chan int,1) 
	var queue chan string = make(chan string)
	var search_pattern string = "Go"

	//WHEN GOROUTINE IS BLOCKING IT RELAXES SYSTEM THREAD AND BACKES THREAD TO SHEDULE
	//GOROUTINE CAN CHANGE SYSTEM THREAD AFTER BLOCKING 
	for i := 0; i < numCPU; i++ {
		//chain var copy by REF
		go func(queue chan string,i int) {
			//REDUCE CPU CASH FLUSH BY FORBID FOR GOROUTINE TO EXECUTE IN ANOTHER THREAD
			runtime.LockOSThread()
			
			a := i+1
			
			//show_thread_id(a)
			
			//block until data is received 
			//block <- 2
			
			//show_thread_id(a)
			
			//until chan close
			for r := range queue { 
				fmt.Printf("GOROUTINE %d parse url: count of " + search_pattern + " match in response is:%d\n",a,strings.Count(r, search_pattern))
			}
			
			block <- 2
		}(queue,i)
	}

	//show_thread_id(0)
	//time.Sleep(2 * time.Second)

	//will block until there are data in the channel
	/*<-block
	<-block
	<-block
	<-block
	show_thread_id(0)
	time.Sleep(2 * time.Second)*/
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//fmt.Printf("scan " + scanner.Text()+"\n")
		if scanner.Text() == "exit" {
			break
		} else {
			queue <- scanner.Text()
		}
	}

	//break for queue cycles in each goroutine
	close(queue)
	<-block
	<-block
	<-block
	<-block
	//val = <-c
	///fmt.Printf("val: %d\n",val)
	
	//fmt.Printf("numCPU: %v\n",numCPU)
}
