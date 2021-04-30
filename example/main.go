package main

import (
	"fmt"

	"github.com/palano/stacktracer"
)

func main() {
	stack := stacktracer.Caller(0)
	fmt.Print(stack.String() + "\n")

	stacks := stacktracer.Callers(0)
	fmt.Print(stacks.String() + "\n")
}
