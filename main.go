package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mikael-kolanowski/abacus/engine"
)

func main() {
	engine := engine.New()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("> ")
		if scanner.Scan() {
			line := scanner.Text()
			result, err := engine.Eval(line)
			if err != nil {
				println("error: ", err)
			} else {
				fmt.Printf("%.3f\n", result.Value)
			}
		} else {
			break
		}
	}
}
