package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main(){
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables
without closure, for two or more funcs to have
that variable would need to be package scope
*/