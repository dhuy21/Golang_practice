package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}

//Example: goimports -w Ex3.go && go build Ex3.go && ./Ex3 1528 758489 890 268 77298
//1528 758489 890 268 77298
//[1528 758489 890 268 77298]
