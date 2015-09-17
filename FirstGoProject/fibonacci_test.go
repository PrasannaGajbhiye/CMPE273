package main

import "testing"

//struct for test pairs
type testPair struct{
	number int64
	fibOfNumber int64
}

// test cases
var tests = []testPair {
	{0,0},
	{1,1},
	{2,1},
	{3,2},
	{4,3},
	{5,5},
	{6,8},
	{7,13},
	{8,21},
	{9,34},
	{10,55},
	//{49,7778742049},
}

// Function to test fib(n) function
func TestFibonacci(t *testing.T) {
	for _, pair:=range tests{
		v := fib(pair.number)
		if v != pair.fibOfNumber {
	           t.Error(
					"For ", pair.number,
					"expected ",pair.fibOfNumber,
					"got ",v,
			)
    	}
	}
}