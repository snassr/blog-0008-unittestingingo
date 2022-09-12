package singlylinkedlist

// SinglyLinkedList implements a list of items.
type SinglyLinkedList interface {
	Append(val int)
	Values() []int
}

// NewSinglyLinked list return a new list.
func NewSinglyLinkedList(vals ...int) SinglyLinkedList {
	return newList(vals...) // swap different implementation constructor here.
}
