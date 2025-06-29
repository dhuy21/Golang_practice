package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, files := range counts {
		for file, n := range files {
			//We use 2 two loop for because we have 2 nested map
			if n > 1 {
				fmt.Printf("%d\t%s\t==>\t%s\n", n, line, file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			//We must intialize the sub-nested map counts[input.Text()] because the initalization in the main function
			// counts := make(map[string]map[string]int) only initialized the map parents counts, doesn't initialize the sub-map
			//Moreover, It's impossible to initialize the sub-map in the func main because we don't know the key input.Text() of the parent-map
			counts[input.Text()] = make(map[string]int)
			//Don't use the := but use = because the map counts have already initialized,
			//so counts[input.Text()] is the value of this map
		}
		counts[input.Text()][f.Name()]++
	}

}

//goimports -w Ex1.4.go && go build Ex1.4.go && ./Ex1.4.go  test1.txt  test2.txt input.txt 2> error.log
