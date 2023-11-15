package stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type person struct {
	name string
	age  int
}

func TestStream_Filter(t *testing.T) {
	numbers := []int{1, 2, 3}
	filteredNumbers := Of[int](numbers).
		Filter(func(n int) bool { return n > 1 }).
		Slice()
	assert.Equal(t, 2, filteredNumbers[0])
	assert.Equal(t, 3, filteredNumbers[1])
	assert.Equal(t, 2, len(filteredNumbers))
}

func TestStructFilter(t *testing.T) {
	persons := []person{{"Bob", 50}, {"Bob", 100}, {"Alice", 30}}
	filteredNumbers := Of[person](persons).Filter(func(p person) bool { return p.name == "Bob" }).Slice()
	assert.Equal(t, person{"Bob", 50}, filteredNumbers[0])
	assert.Equal(t, person{"Bob", 100}, filteredNumbers[1])
	assert.Equal(t, 2, len(filteredNumbers))

	filteredNumbers = Of[person](persons).Filter(func(p person) bool {
		return p.name == "Bob" && p.age == 100
	}).Slice()
	assert.Equal(t, person{"Bob", 100}, filteredNumbers[0])
	assert.Equal(t, 1, len(filteredNumbers))
}

func TestStream_Map(t *testing.T) {
	persons := []person{{"Bob", 50}, {"Bob", 100}, {"Alice", 30}}
	names := Map[person, string](persons, func(p person) string {
		return p.name
	}).Filter(func(name string) bool { return name != "Bob" }).
		Slice()

	assert.Equal(t, "Alice", names[0])
	assert.Equal(t, 1, len(names))
}

func TestStream_Sort(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}
	sorted := Of(numbers).Sort(func(a, b int) bool { return a < b }).Slice()
	assert.Equal(t, []int{1, 2, 3, 4, 5}, sorted)

	persons := []person{{"Bob", 100}, {"Bob", 50}, {"Alice", 30}}
	sortedPersons := Of[person](persons).Sort(func(a, b person) bool {
		if a.name == b.name {
			return a.age < b.age
		}
		return a.name < b.name
	}).Slice()

	assert.Equal(t, []person{{"Alice", 30}, {"Bob", 50}, {"Bob", 100}}, sortedPersons)
}

func TestStream_AnyMatch(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}
	match := Of(numbers).AnyMatch(func(a int) bool { return a == 1 })
	assert.Equal(t, true, match)

	match = Of(numbers).AnyMatch(func(a int) bool { return a == 6 })
	assert.Equal(t, false, match)
}

func TestStream_NoneMatch(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}
	match := Of(numbers).NoneMatch(func(a int) bool { return a == 1 })
	assert.Equal(t, false, match)

	match = Of(numbers).NoneMatch(func(a int) bool { return a == 6 })
	assert.Equal(t, true, match)
}

func TestStream_FindFirst(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}
	first := Of(numbers).FindFirst()
	assert.Equal(t, 5, first.Get())

	first = Of(numbers).Filter(func(i int) bool { return i > 5 }).FindFirst()
	assert.Equal(t, false, first.IsPresent())
}

func TestFlatMap(t *testing.T) {
	numbers := [][]int{{1, 2, 3}, {4, 5, 6}}
	flattened := FlatMap[int](numbers).Slice()
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, flattened)
}

func TestGroupBy(t *testing.T) {
	persons := []person{{"Bob", 100}, {"Bob", 50}, {"Alice", 30}}
	group := GroupBy[string, person](persons, func(p person) string { return p.name })
	assert.Equal(t, 2, len(group["Bob"]))
	assert.Equal(t, 1, len(group["Alice"]))
}
