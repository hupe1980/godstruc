package queue

type Queue interface {
	Push(elem interface{})
	Pop() (interface{}, bool)
	Len() int
	Empty() bool
	Front() (interface{}, bool)
}
