package leetcode

func isValid(s string) bool {
	m := make(map[string]string, 3)
	m[")"] = "("
	m["}"] = "{"
	m["]"] = "["

	stack := Stack{}

	for _, ch := range s {
		if isClosing(string(ch)) {
			if stack.Len() == 0 {
				return false
			}

			if m[string(ch)] != stack.Pop() {
				return false
			}
		} else {
			stack.Push(string(ch))
		}
	}

	if stack.Len() != 0 {
		return false
	}

	return true
}

func isClosing(s string) bool {
	if s == ")" || s == "}" || s == "]" {
		return true
	}

	return false
}

type Stack struct {
	lifo []string
}

func (s *Stack) Push(val string) {
	s.lifo = append(s.lifo, val)
}

func (s *Stack) Pop() string {
	el := s.lifo[len(s.lifo)-1]
	s.lifo = s.lifo[:len(s.lifo)-1]

	return el
}

func (s *Stack) Len() int {
	return len(s.lifo)
}
