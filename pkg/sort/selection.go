package sort

import "github.com/Felyp-Henrique/algorithm/pkg/types"

type SelectionSort[T any] struct {
	list []types.Comparable[T]
}

func (s *SelectionSort[T]) SetList(list []types.Comparable[T]) {
	s.list = list
}

func (s *SelectionSort[T]) GetList() []types.Comparable[T] {
	return s.list
}

func (s *SelectionSort[T]) Sort() []T {
	var (
		listNew        []T                   = make([]T, len(s.list))
		listOld        []types.Comparable[T] = s.list
		elementIndex   int
		elementGreater T
	)
	for len(listOld) > 0 {
		elementIndex = -1
		elementGreater = *new(T)
		for index, value := range s.list {
			if value.IsGreater(elementGreater) {
				elementIndex = index
				elementGreater = value.GetValue()
			}
		}
		listNew = append(listNew, elementGreater)
		listOld = append(listOld[:elementIndex], listOld[elementIndex+1:]...)
	}
	return listNew
}
