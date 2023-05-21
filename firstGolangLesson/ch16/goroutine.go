package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tick tick ", times)
		time.Sleep(time.Second)
	}
}

func main() {
	go running()

	var input string
	fmt.Scanln(&input)

	fmt.Println(input)
}
