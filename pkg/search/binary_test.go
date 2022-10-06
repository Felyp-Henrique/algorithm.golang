package search_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Felyp-Henrique/algorithm/pkg/search"
	"github.com/go-faker/faker/v4"
)

func TestBinarySearchFind(t *testing.T) {
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	index := random.Intn(1000)
	peoples := newpeoples()
	people := peoples[index]
	search := search.NewBinarySearch(peoples)
	position := search.Find(people)
	if position != index {
		t.Fail()
	}
}

func BenchmarkBinarySearchFind(b *testing.B) {
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	index := random.Intn(1000)
	peoples := newpeoples()
	people := peoples[index]
	search := search.NewBinarySearch(peoples)
	b.StartTimer()
	position := search.Find(people)
	b.StopTimer()
	if position != index {
		b.Fail()
	}
}

type people struct {
	id       int
	username string
}

func newpeople(id int) people {
	return people{
		id:       id,
		username: faker.Username(),
	}
}

func (u people) Index() int {
	return u.id
}

func newpeoples() []people {
	var (
		i       int
		peoples []people = []people{}
	)
	for i = 0; i < 1000; i++ {
		peoples = append(peoples, newpeople(i))
	}
	return peoples
}
