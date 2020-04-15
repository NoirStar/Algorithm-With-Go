package dataStruct

type Stack struct {
	li *DoubleLinkedList
}

func NewStack() *Stack {
	// 새로운 스택 초기화 함수
	return &Stack{li: &DoubleLinkedList{}}
}

func (s *Stack) Push(val int) {
	s.li.AddNode(val)
}

func (s *Stack) Pop(val int) int {
	back := s.li.Back()
	s.li.PopBack()
	return back
}
