package main

import (
		"testing"
		 "time" 
		"fmt"
		)

//struct for test pairs
type testPair struct{
	sleepDur int
	actualSleepDur int
	chann chan string
}

// test cases
var tests = []testPair {
	{1,1,make(chan string)},	
}

// Function to test fib(n) function
func TestSleep(t *testing.T) {
	for _, pair:=range tests{
		t0:=time.Now()
		time.Sleep(time.Second*time.Duration(pair.sleepDur))
		t1:=time.Now()
		if(mySleep(pair.chann, pair.sleepDur) - t1.Sub(t0)>1){
			
			t.Error(
					"For ", pair.sleepDur,
					"expected ", t1.Sub(t0),
					"got :", mySleep(pair.chann, pair.sleepDur),
					)
		}
		
	}
}