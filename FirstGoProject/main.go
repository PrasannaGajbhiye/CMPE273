// Fibonacci project main.go
package main

import (
	"fmt" //format package
)

//Entry Function
func main() {
	var number int64
	var x int64
	
	//Prompt a number & scan the same
	fmt.Print("Enter a number,n to find fib(n) :")
	fmt.Scanf("%d",&number)
	
	//Invoke fibonacci function
	x = fib(number)
	
	fmt.Println("fib(n) is :",x)
}

//Function to find fibonacci sequence
func fib(num int64) int64{
	
	var retVal int64

	if num == 0 {
		retVal=0	//fib(0)=0
	}else if num == 1 {
		retVal=1	//fib(1)=1
	}else{
		retVal= fib(num-2)+ fib(num-1)
	}
	return retVal
}