package types

type Equals[T any] interface {
	IsEquals(other T) bool
}

type Greater[T any] interface {
	IsGreater(other T) bool
}

type Smaller[T any] interface {
	IsSmaller(other T) bool
}

type Comparable[T any] interface {
	Equals[T]
	Greater[T]
	Smaller[T]

	GetValue() T
}
