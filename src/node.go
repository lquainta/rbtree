package src

type Node[T any] struct {
	data   T
	color  bool // true = red, false = black
	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
	count  int
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{data: data, count: 1}
}
