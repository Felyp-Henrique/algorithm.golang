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
		elementGreater T
	)
	for len(listOld) > 0 {
		for index, value := range s.list {
			if value != nil && value.IsGreater(elementGreater) {
				elementGreater = value.GetValue()
				listOld[index] = nil
			}
		}
		listNew = append(listNew, elementGreater)
	}
	return listNew
}
