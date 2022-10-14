package list

type LinkedNode[D any] struct {
	data     D
	next     *LinkedNode[D]
	previous *LinkedNode[D]
}

func (l LinkedNode[D]) Data() D {
	return l.data
}

func (l LinkedNode[D]) Next() *LinkedNode[D] {
	return l.next
}

func (l LinkedNode[D]) Previous() *LinkedNode[D] {
	return l.previous
}

type LinkedList[D any] struct {
	length int
	first  *LinkedNode[D]
	last   *LinkedNode[D]
}

func (l LinkedList[D]) Length() int {
	return l.length
}

func (l LinkedList[D]) First() *LinkedNode[D] {
	return l.first
}

func (l LinkedList[D]) Last() *LinkedNode[D] {
	return l.last
}

func (l LinkedList[D]) Get(index int) *LinkedNode[D] {
	return nil
}

func (l *LinkedList[D]) Add(node *LinkedNode[D]) {
	l.length += 1
}

func (l *LinkedList[D]) Remove(index int) {
	l.length -= 1
}
