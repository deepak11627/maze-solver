package maze

import (
	"fmt"
)

// Maze struct represents a maze of x * y dimension
type Maze struct {
	cells  [][]Cell
	width  int
	height int
}

// Option allow to override default properties of Maze
type Option func(m *Maze)

// SetMazeCells allows to Initialise the maze
func SetMazeCells(c [][]Cell) func(m *Maze) {
	return func(m *Maze) {
		m.cells = c
	}
}

// NewMaze return a pointer to a Maze struct
func NewMaze(opts ...Option) *Maze {
	m := &Maze{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// FindPath find the paht of the maze
func (m *Maze) FindPath() (*Path, error) {
	fmt.Println("finding path")
	return &Path{}, nil
}

// Display renders the Maze on screen
func (m *Maze) Display() {
	for _, cells := range m.cells {
		for _, cell := range cells {
			fmt.Print(cell.char)
		}
		fmt.Println()
	}
}
