package main

import (
		"testing"
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
	{2,2,make(chan string)},
	{3,3,make(chan string)},	
}

// Function to test fib(n) function
func TestSleep(t *testing.T) {
	for _, pair:=range tests{
		go mySleep(pair.chann, pair.sleepDur)
	}
}