package search

import "github.com/Felyp-Henrique/algorithm/pkg/interfaces"

type BinarySearch[T interfaces.Indexable] struct {
	values []T
	low    int
	high   int
}

func NewBinarySearch[T interfaces.Indexable](values []T) BinarySearch[T] {
	return BinarySearch[T]{
		values: values,
		low:    0,
		high:   len(values) - 1,
	}
}

func (b BinarySearch[T]) Low() int {
	return b.low
}

func (b BinarySearch[T]) High() int {
	return b.high
}

func (b BinarySearch[T]) Middle() int {
	return (b.low + b.high) / 2
}

func (b BinarySearch[T]) Find(other T) int {
	var (
		value T
	)
	for ; b.low <= b.high; value = b.values[b.Middle()] {
		if value.Index() == other.Index() {
			return b.Middle()
		}
		if value.Index() < other.Index() {
			b.low += 1
		}
		if value.Index() > other.Index() {
			b.high -= 1
		}
	}
	return -1
}
