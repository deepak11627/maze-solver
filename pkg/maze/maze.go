package maze

import (
	"errors"
	"fmt"
	"strconv"
)

// Char a type alias for string to confusion in signle and multi character strings
type Char string

type Config struct {
	startChar Char
	endChar   Char
	wallChar  Char
	openChar  Char
}

func NewMazeConfig(s, e, w, o Char) *Config {
	return &Config{s, e, w, o}
}

// Maze struct represents a maze of x * y dimension
type Maze struct {
	cells  [][]Cell
	width  int
	height int

	startAt Cell
	endAt   Cell

	config *Config
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
func NewMaze(c *Config, opts ...Option) *Maze {
	m := &Maze{config: c}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// FindPath find the path of the maze
// Algorithm mentioned at https://www.cs.bu.edu/teaching/alg/maze/ has been followed to find the path
func (m *Maze) FindPath() (*Path, error) {
	p := &Path{}
	if m.traverse(m.startAt.x, m.startAt.y, p) == false {
		return p, errors.New("no path found")
	}
	return p, nil
}

func (m *Maze) traverse(x, y int, p *Path) bool {
	// 	if (x,y outside maze) return false
	if m.isOutsideMaze(x, y) {
		return false
	}

	// if (x,y is goal) return true
	if m.isMazeGoal(x, y) {
		return true
	}
	// if (x,y not open) return false
	if m.isOpen(x, y, p) {
		return false
	}

	// mark x,y as part of solution path
	p.Push(strconv.Itoa(x) + "-" + strconv.Itoa(y))
	// if (FIND-PATH(North of x,y) == true) return true
	if m.traverse(x, y-1, p) {
		return true
	}

	// if (FIND-PATH(East of x,y) == true) return true
	if m.traverse(x+1, y, p) {
		return true
	}

	// if (FIND-PATH(South of x,y) == true) return true
	if m.traverse(x, y+1, p) {
		return true
	}

	// if (FIND-PATH(West of x,y) == true) return true
	if m.traverse(x-1, y, p) {
		return true
	}
	// unmark x,y as part of solution path
	p.Pop()
	// return false
	return false
}

func (m *Maze) isOutsideMaze(x, y int) bool {
	return !(x >= 0 && x < m.width && y >= 0 && y < m.height)
}

func (m *Maze) isMazeGoal(x, y int) bool {
	return x == m.endAt.x && y == m.endAt.y
}

func (m *Maze) isOpen(x, y int, p *Path) bool {
	return m.cells[x][y].char == m.config.openChar && !p.Exists(strconv.Itoa(x)+"-"+strconv.Itoa(y))
}

// Display renders the Maze on screen
func (m *Maze) Display() {
	fmt.Println("Here is the maze read from the file")
	for _, cells := range m.cells {
		for _, cell := range cells {
			fmt.Print(cell.char)
		}
		fmt.Println()
	}
}

func (m *Maze) findStartPonit() (Cell, error) {
	// TODO find start point
	return Cell{}, nil
}

func (m *Maze) findEndPonit() (Cell, error) {
	// TODO find end point
	return Cell{}, nil
}

func (m *Maze) isValidChar(ch Char) bool {
	validChars := map[Char]struct{}{m.config.wallChar: struct{}{}, m.config.startChar: struct{}{}, m.config.openChar: struct{}{}, m.config.endChar: struct{}{}}
	if _, ok := validChars[ch]; ok {
		return true
	}
	return false
}

func (m *Maze) isStartingChar(char Char) bool {
	if char == m.config.startChar {
		return true
	}
	return false
}

func (m *Maze) isEndingChar(char Char) bool {
	if char == m.config.endChar {
		return true
	}
	return false
}

func (m *Maze) isWallChar(char Char) bool {
	if char == m.config.wallChar {
		return true
	}
	return false
}

func (m *Maze) isOpenChar(char Char) bool {
	if char == m.config.openChar {
		return true
	}
	return false
}
