package maze

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

//type MazeFileErr error

var InvalidMazeDimensionErr = errors.New("Invalid arguments for maze dimension")
var InvalidMazeRowErr = errors.New("the row %d is invalid in the maze as the number of cell don't match with maze dimension")
var InvalidMazeCharacterErr = errors.New("Invalid maze character. Only #, e, o and x are the only allowed characters.")

// MazeReader reads maze from a source and create a maze
type MazeReader struct {
	source io.Reader
}

// NewMazeReader return a pointer to a Maze struct
func NewMazeReader(r io.Reader) *MazeReader {
	return &MazeReader{source: r}
}

func (mr *MazeReader) Read(m *Maze) error {
	scanner := bufio.NewScanner(mr.source)
	isFirstLine := true
	for scanner.Scan() {
		if isFirstLine {
			s := strings.Split(scanner.Text(), " ")
			if len(s) != 2 {
				return InvalidMazeDimensionErr
			}
			height, err := strconv.Atoi(s[0])
			if err != nil {
				panic(err)
			}
			width, err := strconv.Atoi(s[1])
			if err != nil {
				panic(err)
			}
			m.width = width
			m.height = height

			isFirstLine = false
			continue
		}
		row := scanner.Text()
		if len(row) != m.width {
			return InvalidMazeRowErr
		}

		mazeRow, err := mr.readLine(row)
		if err != nil {
			return err
		}
		m.cells = append(m.cells, mazeRow)
	}
	return nil
}

func (mr *MazeReader) readLine(row string) ([]Cell, error) {
	var mazeRow []Cell
	for _, ch := range row {
		if !mr.isValidChar(string(ch)) {
			return mazeRow, InvalidMazeCharacterErr
		}
		mazeRow = append(mazeRow, Cell{char: string(ch)})
	}
	return mazeRow, nil
}

func (mr *MazeReader) isValidChar(ch string) bool {
	validChars := map[string]struct{}{"#": struct{}{}, "e": struct{}{}, "o": struct{}{}, "x": struct{}{}}
	if _, ok := validChars[ch]; ok {
		return true
	}
	return false
}
