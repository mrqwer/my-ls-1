package main

import (
	"my-ls-1/internal/helper"
	"my-ls-1/internal/ls"
	"os"
)

func main() {
	parameters := os.Args[1:]
	if len(parameters) == 0 {
		ls.DefaultCase()
		return
	}
	splitParams := helper.SplitOptionsAndFilenames(parameters)
	ls.Program(splitParams)
}
