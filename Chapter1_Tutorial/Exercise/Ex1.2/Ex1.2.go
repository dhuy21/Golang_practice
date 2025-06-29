package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Index  Value")
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, ": ", os.Args[i])
	}
}

/*goimports -w Ex2.go && go build Ex2.go && ./Ex2 1528 758489 890 268 77298
Index  Value
0 :  ./Ex2
1 :  1528
2 :  758489
3 :  890
4 :  268
5 :  77298*/
