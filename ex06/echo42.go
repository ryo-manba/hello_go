package main

import (
	"flag"
	"fmt"
)

func Separate(sep string, Args []string) {
	if len(Args) == 0 || sep == "-n" {
		return
	}

	for i := 0; i < len(Args); i++ {
		fmt.Printf("%s", Args[i])
		if i < len(Args)-1 {
			fmt.Printf("%s", sep)
		}
	}
	fmt.Printf("\n")
}

func OmitNewLine(Args []string) {
	for i := 0; i < len(Args); i++ {
		fmt.Printf("%s", Args[i])
		if i < len(Args)-1 {
			fmt.Printf(" ")
		}
	}
}

func main() {
	s := flag.String("s", " ", "separator")
	n := flag.Bool("n", false, "omit trailing newline")
	flag.Parse()

	if *s != " " && *n == true {
		return
	}

	if *s != " " {
		Separate(*s, flag.Args())
		return
	}

	if *n == true {
		OmitNewLine(flag.Args())
		return
	}

	for i := 0; i < len(flag.Args()); i++ {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%s", flag.Arg(i))
	}
	fmt.Printf("\n")
}
