package maze

import (
	"fmt"
)

// Path implimented using stack
type Path []Cell

// Empty deletes all path elements
func (s Path) Empty() bool { return len(s) == 0 }

// Push adds a step to the path
func (s *Path) Push(v Cell) { (*s) = append((*s), v) }

// Pop remove a step from the path
func (s *Path) Pop() Cell {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func (s *Path) isPreviousStep(e Cell) bool {
	if len(*s) == 0 {
		return false
	}
	val := s.Pop()
	s.Push(val)
	if val == e {
		return true
	}
	return false
}

// Exists checks if an element exists
func (s *Path) Exists(e Cell) bool {
	for _, p := range *s {
		if p == e {
			if s.isPreviousStep(e) {
				return false
			}
			return true
		}
	}
	return false
}

// Traverse renders the path
func (s *Path) Traverse() string {
	var path []byte
	for _, p := range *s {
		path = append(path, fmt.Sprintf("%d,%d -> ", p.x, p.y)...)
	}
	return string(path)
}
