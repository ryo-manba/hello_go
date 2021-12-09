package main

import (
	"fmt"
	"os"
	"regexp"
)

func CheckRegexp(reg, str string) bool {
	return (regexp.MustCompile(reg).MatchString(str))
}

func CheckAddress(addr string) bool {
	if len(addr) > 256 {
		return false
	}
	return (CheckRegexp(`^[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+$`, addr))
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Invalid argument\n")
		return
	}
	for i := 1; i < len(os.Args); i++ {
		addr := os.Args[i]
		if CheckAddress(addr) {
			fmt.Printf("%s is a valid email address.\n", addr)
		} else {
			fmt.Printf("%s is NOT a valid email address.\n", addr)
		}
	}

}
