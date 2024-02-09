package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNumbers(n int) {
	for i := 0; i < 500; i++ {
		fmt.Print(n)
	}
}

func parallelFunction() {
	fmt.Println("\nParallel Code Execution:")
	runtime.GOMAXPROCS(3)
	go printNumbers(0)
	go printNumbers(1)
	time.Sleep(time.Second)
}

func concurrentFunction() {
	fmt.Println("Concurrent Code Execution:")
	runtime.GOMAXPROCS(1)
	go printNumbers(0)
	go printNumbers(1)
	time.Sleep(time.Second)
}

func main() {
	concurrentFunction()
	parallelFunction()
}
