package main

import (
	"fmt"
	"os"
	"strings"
)

func version1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func version2() {
	fmt.Println(os.Args[0:])
}

func version3() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	version1()
	version2()
	version3()
}
