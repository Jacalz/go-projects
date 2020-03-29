// Calculator using logic from https://thorstenball.com/blog/2016/11/16/putting-eval-in-go/ in order to get evaluation of mathermatical functions.

package main

import (
	"bufio"
	"eval"
	"fmt"
	"go/parser"
	"log"
	"os"
)

func scan() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		os.Exit(2)
	}

	line := scanner.Text()
	exp, err := parser.ParseExpr(line)
	if err != nil {
		log.Fatalf("parsing failed: %s\n", err)
	}

	return eval.Eval(exp)

}

func main() {
	fmt.Print("Enter mathematical expression: ")
	ans := scan()

	fmt.Println("Result:", ans)

}
