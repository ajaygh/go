package main

import (
	"fmt"
)

//See power of self-referential functions
type option func(*Student) option

type Student struct {
	FirstName, LastName string
	Verbosity           int
}

func (s *Student) Option(opts ...option) (prev option) {
	for _, op := range opts {
		prev = op(s)
	}
	return prev
}
func (s *Student) Print() {
	switch s.Verbosity {
	case 1:
		fmt.Println("First Name is ", s.FirstName)
	case 2:
		fmt.Println("Last name is", s.LastName)
	case 3:
		fmt.Println("Full name is", s.FirstName+" "+s.LastName)
	}
}

func setVerbosity(v int) option {
	return func(s *Student) option {
		prev := s.Verbosity
		s.Verbosity = v
		return setVerbosity(prev)
	}
}

func setFirstName(f string) option {
	return func(s *Student) option {
		prevName := s.FirstName
		s.FirstName = f
		return setFirstName(prevName)
	}
}

func main() {
	s := &Student{"Sachin", "Tendulkar", 1}
	fmt.Println("Initial ", s)
	s.Option(setVerbosity(2))
	s.Print()
	prev := s.Option(setVerbosity(3), setFirstName("shyam"))
	s.Print()
	s.Option(prev)
	s.Print()
}
