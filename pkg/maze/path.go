package maze

import "fmt"

// Path implimented using stack
type Path []string

// Empty deletes all path elements
func (s Path) Empty() bool { return len(s) == 0 }

// Push adds a step to the path
func (s *Path) Push(v string) { (*s) = append((*s), v) }

// Pop remove a step from the path
func (s *Path) Pop() string {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

// Exists checks if an element exists
func (s *Path) Exists(e string) bool {
	for _, p := range *s {
		if p == e {
			return true
		}
	}
	return false
}

// Traverse renders the path
func (s *Path) Traverse() {
	for _, p := range *s {
		fmt.Println(p + " ")
	}
}
