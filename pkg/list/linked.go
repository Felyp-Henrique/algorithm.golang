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
	var (
		node    *LinkedNode[T] = l.first
		current int            = 0
	)
	if l.length == 0 || index >= l.length || index < 0 {
		return nil
	}
	if index == 0 {
		return l.first
	}
	if index == l.length-1 {
		return l.last
	}
	for current != index {
		node = node.next
		current += 1
	}
	return node
}

func (l *LinkedList[T]) Add(value T) {
	var (
		node *LinkedNode[T] = &LinkedNode[T]{
			value: value,
		}
	)
	if l.first == nil {
		l.first = node
	}
	if l.last == nil {
		l.last = node
		node.previous = l.last
	} else {
		l.last.next = node
		node.previous = l.last
		l.last = node
	}
	l.length += 1
}

func (l *LinkedList[T]) Remove(index int) {
	var (
		node *LinkedNode[T] = nil
	)
	if node = l.Get(index); node != nil {
		if node == l.last {
			//l.last.previous
		}
		node.next = node.previous
		node.previous = node.next
		node = nil
		l.length -= 1
	}
}

func (l *LinkedList[T]) Push(index int, value T) {
	l.length += 1
}

func (l *LinkedList[T]) Pop(index int) *LinkedNode[T] {
	l.length -= 1
	return l.Get(index)
}
