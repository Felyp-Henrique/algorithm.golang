package list

type LinkedNode[T any] struct {
	value    T
	next     *LinkedNode[T]
	previous *LinkedNode[T]
}

func (l LinkedNode[T]) Data() T {
	return l.value
}

func (l LinkedNode[T]) Next() *LinkedNode[T] {
	return l.next
}

func (l LinkedNode[T]) Previous() *LinkedNode[T] {
	return l.previous
}

type LinkedList[T any] struct {
	length int
	first  *LinkedNode[T]
	last   *LinkedNode[T]
}

func (l LinkedList[T]) Length() int {
	return l.length
}

func (l LinkedList[T]) First() *LinkedNode[T] {
	return l.first
}

func (l LinkedList[T]) Last() *LinkedNode[T] {
	return l.last
}

func (l LinkedList[T]) Get(index int) *LinkedNode[T] {
	return nil
}

func (l *LinkedList[T]) Add(value T) {
	var (
		node *LinkedNode[T] = &LinkedNode[T]{}
	)
	node.value = value
	l.length += 1
}

func (l *LinkedList[T]) Remove(index int) {
	l.length -= 1
}

func (l *LinkedList[T]) Push(index int, value T) {
	l.length += 1
}

func (l *LinkedList[T]) Pop(index int) *LinkedNode[T] {
	l.length -= 1
	return l.Get(index)
}
