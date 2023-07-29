package main

import (
	"fmt"
	"lemin/lib"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		numberOfAnts, antFarm, fileContent, shouldReturn := lib.ParseFile(args[1])
		if shouldReturn {
			return
		}
		fmt.Println(fileContent)
		paths, err := lib.FindPaths(antFarm)

		if err != nil {
			fmt.Println("No path found")
			return
		}

		lib.MoveAnts(paths, numberOfAnts)
	} else {
		fmt.Println("[USAGE] : ./lem-in path/to/your/file.txt")
	}
}
