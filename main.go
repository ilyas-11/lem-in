package main

import (
	"fmt"

	"os"
	"strings"

	"lem-in/parser"
	"lem-in/pathfinder"
	"lem-in/simulator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage error: program requires exactly one argument - the path to input file")
		fmt.Println("example: ./lem-in input.txt")
		return
	}
	if !(strings.HasSuffix(os.Args[1], ".txt")) {
		fmt.Println("file format error: input file must be a .txt file")
		fmt.Println("example: input.txt, farm.txt")
		return
	}
	data, err := parser.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	Paths, err := pathfinder.Grouppaths(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if Paths.Pathgroup == nil {
		fmt.Println("No paths found from start to end.")
		return
	}
	fmt.Println(Paths.Pathgroup)
	

	simulator.Simulation(data, Paths)
	simulator.Printpaths(data, Paths)
}
