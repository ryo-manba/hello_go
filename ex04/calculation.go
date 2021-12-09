package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func calculationStr(Args []string) (string, bool) {
	if len(Args) != 3 {
		return "", false
	}
	a, A := strconv.Atoi(Args[1])
	b, B := strconv.Atoi(Args[2])
	if A != nil || B != nil || b == 0 {
		return "", false
	}
	sum := a + b
	diff := a - b
	mul := a * b
	div := a / b
	s := fmt.Sprintf("sum: %d\ndifference: %d\nproduct: %d\nquotient: %d\n",
		sum, diff, mul, div)
	return s, true
}

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
