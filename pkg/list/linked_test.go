package list_test

import (
	"fmt"
	"testing"

	"github.com/Felyp-Henrique/algorithm/pkg/list"
)

func TestLinkedList(t *testing.T) {
	var (
		list list.LinkedList[int] = list.LinkedList[int]{}
	)
	list.Add(20)
	list.Add(80)
	list.Add(50)
	list.Add(70)
	fmt.Println(list.Get(1).Data())
	fmt.Println(list.Get(2).Data())
	fmt.Println(list.Get(3).Data())
	fmt.Println(list.Length())
	list.Remove(4)
	fmt.Println(list.Get(1).Data())
	fmt.Println(list.Get(2).Data())
	fmt.Println(list.Get(3))
	fmt.Println(list.Length())
}
