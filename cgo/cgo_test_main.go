package main

/*
#include <stdlib.h>

void emptyCFunction() {
	// This is an empty C function
}
*/
import "C"
import (
	"fmt"
	"time"
)

func emptyFunction() {
	// This is an empty function
}
func main() {
	const iterations = 10000000
	fmt.Printf("iterations: %v\n", iterations)
	start := time.Now()
	for i := 0; i < iterations; i++ {
		emptyFunction()
	}
	duration := time.Since(start)
	fmt.Printf("Go function call took %s, avg: %v\n", duration, duration/iterations)
	start = time.Now()
	for i := 0; i < iterations; i++ {
		C.emptyCFunction()
	}
	duration = time.Since(start)
	fmt.Printf("C function call took %s, avg: %v\n", duration, duration/iterations)
}
