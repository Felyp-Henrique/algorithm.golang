package search_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Felyp-Henrique/algorithm/pkg/search"
	"github.com/go-faker/faker/v4"
)

var (
	peoples []people
	count   int
)

func TestMain(m *testing.M) {
	count = 1_000_000
	peoples = newpeoples()
	m.Run()
}

func TestBinarySearchFind(t *testing.T) {
	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	index := random.Intn(count)
	people := peoples[index]
	search := search.NewBinarySearch(peoples)
	position := search.Find(people)
	if position != index {
		t.Fail()
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
	for i = 0; i < count; i++ {
		peoples = append(peoples, newpeople(i))
	}
	return peoples
}
