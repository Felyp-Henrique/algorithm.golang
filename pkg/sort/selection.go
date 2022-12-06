package sort

import "github.com/Felyp-Henrique/algorithm/pkg/types"

type SelectionSort[T types.Indexable] struct {
	list []T
}

func (s *SelectionSort[T]) SetList(list []T) {
	s.list = list
}

func (s *SelectionSort[T]) GetList() []T {
	return s.list
}

func (s *SelectionSort[T]) Sort() []T {
	var (
		result []T = make([]T, len(s.list))
		larger int = -99
	)
	for index, value := range s.list {
		if value.Index() > larger {
			result[index] = value;
			larger = value.Index()
		}
	}
	return []T{}
}
