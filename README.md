Recreated useful java stream functionality.

# Examples
The following struct was used to demonstrate functionalities upon struct types.
```go
type person struct {
    name string
    age  int
}
```
## Filtering
```go
numbers := []int{1, 2, 3}
filteredNumbers := Of[int](numbers).Filter(func(n int) bool { return n > 1 }).Slice()
```

```go
persons := []person{{"Bob", 50}, {"Bob", 100}, {"Alice", 30}}
filteredNumbers := Of[person](persons).Filter(func(p person) bool { return p.name == "Bob" }).Slice()
```

## Mapping
```go
persons := []person{{"Bob", 50}, {"Bob", 100}, {"Alice", 30}}
names := Map[person, string](persons, func(p person) string {
    return p.name
}).Filter(func(name string) bool { return name != "Bob" }).
    Slice()
```

## Sorting
```go
numbers := []int{5, 4, 3, 2, 1}
sorted := Of(numbers).Sort(func(a, b int) bool { return a < b }).Slice()

persons := []person{{"Bob", 100}, {"Bob", 50}, {"Alice", 30}}
sortedPersons := Of[person](persons).Sort(func(a, b person) bool {
    if a.name == b.name {
        return a.age < b.age
    }
    return a.name < b.name
}).Slice()
```

## Any Match
```go
// This will return TRUE
numbers := []int{5, 4, 3, 2, 1}
match := Of(numbers).AnyMatch(func(a int) bool { return a == 1 })

// This will return FALSE
match = Of(numbers).AnyMatch(func(a int) bool { return a == 6 })
assert.Equal(t, false, match)
```

## None Match
```go
// This will return false
numbers := []int{5, 4, 3, 2, 1}
match := Of(numbers).NoneMatch(func(a int) bool { return a == 1 })

// This will return true
match = Of(numbers).NoneMatch(func(a int) bool { return a == 6 })
```