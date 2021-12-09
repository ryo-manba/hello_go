package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	num, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < num; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("*")
			num--
		}
		fmt.Printf("\n")
	}
}
