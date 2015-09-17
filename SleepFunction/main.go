// SleepFunction project main.go
package main

import (
	"fmt"
	"time"
)

var sleepDuration int

// Entry Function
func main(){
	// Channel to send and receive messages
	var c chan string = make(chan string)
	
	// Prompt & scan for the duration of sleep
	fmt.Println("Please enter the duration of sleep in seconds: ")
	fmt.Scanf("%d",&sleepDuration)
	
	// Concurrently invoke pinger & printer functions
	go pinger(c)
	go printer(c)
	
    var input string
    fmt.Scanln(&input)
}

// Pinger Function
func pinger(c chan string){
	for i:=0;;i++{
		c<-"ping"
	}
}

// Printer Function
func printer(c chan string){
	for {
		msg := <-c
		fmt.Println(msg)
		rVal:= mySleep(c,sleepDuration)
		fmt.Println(rVal)
	}
}

// Own sleep function
func mySleep(c chan string, seconds int) time.Duration {
	t0:=time.Now()
	for i:=0 ; i< seconds ; i++{
			<-time.After(time.Second)		
	}
	t1:=time.Now()
	return t1.Sub(t0)
	
}