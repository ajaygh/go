package main

import "container/list"
import "fmt"

func revStack(s *list.List) {
	if s.Len() < 2 {
		return
	}
	top := s.Back()
	revStack(s)
	putAtBottom(s, top)
	fmt.Println("top is ", top, "stack ", s)
}

func putAtBottom(s *list.List, top *list.Element) {
	if s.Len() == 0 {
		s.PushBack(top)
		return
	}
	tmptop := s.Back()
	s.Remove(tmptop)
	putAtBottom(s, top)
	s.PushBack(tmptop)
}

func main() {
	s := *list.New()
	s.PushBack(1)
	s.PushBack(2)
	s.PushBack(3)
	s.PushBack(4)
	s.PushBack(5)
	fmt.Println("Before reverse ", s)
	revStack(&s)
	fmt.Println("After reverse ", s)
}
