package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/deepak11627/maze-solver/pkg/maze"
)

func main() {
	// Flags for the program
	mazeInputFile := flag.String("file", "./sample/maze-sample1.txt", "Path to the maze file to read")
	flag.Parse()

	// Open maze file
	file, err := os.Open(*mazeInputFile)
	failOnError(err)
	defer file.Close()

	// Maze file reader
	mr := maze.NewMazeReader(file)
	m := maze.NewMaze()

	// Read maze file onto Maze struct
	err = mr.Read(m)
	failOnError(err)

	// Display the Maze
	m.Display()

	// Find path of the maze, displays error if not found
	_, err = m.FindPath()
	failOnError(err)
}

func failOnError(err error) {
	if err != nil {
		fmt.Println("Failed to process maze due to error ", err)
		os.Exit(1)
	}
}
