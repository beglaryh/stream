package stream

import . "github.com/beglaryh/optional"

type Stream[T any] struct {
	ts []T
}

func Of[T any](ts []T) Stream[T] {
	return Stream[T]{ts: ts}
}

func (stream Stream[T]) Filter(filter func(t T) bool) Stream[T] {
	ns := Stream[T]{}
	for _, t := range stream.ts {
		if filter(t) {
			ns.ts = append(ns.ts, t)
		}
	}
	return ns
}

func Map[F, T any](fs []F, mapper func(f F) T) Stream[T] {
	ns := Stream[T]{}
	for _, e := range fs {
		nv := mapper(e)
		ns.ts = append(ns.ts, nv)
	}
	return ns
}

func FlatMap[T any](input [][]T) Stream[T] {
	var ts []T
	for _, array := range input {
		for _, t := range array {
			ts = append(ts, t)
		}
	}
	return Stream[T]{ts: ts}
}

func GroupBy[K comparable, T any](ts []T, getKey func(t T) K) map[K][]T {
	response := map[K][]T{}
	for _, t := range ts {
		key := getKey(t)
		response[key] = append(response[key], t)
	}
	return response
}

func (stream Stream[T]) Sort(sortFunction func(a, b T) bool) Stream[T] {
	ns := mergeSort[T](stream.ts, sortFunction)
	return Stream[T]{ts: ns}
}

func (stream Stream[T]) AnyMatch(anyFunction func(t T) bool) bool {
	for _, e := range stream.ts {
		if anyFunction(e) {
			return true
		}
	}
	return false
}

func (stream Stream[T]) NoneMatch(anyFunction func(t T) bool) bool {
	for _, e := range stream.ts {
		if anyFunction(e) {
			return false
		}
	}
	return true
}

func (stream Stream[T]) FindFirst() *Optional[T] {
	if len(stream.ts) == 0 {
		empty := Empty[T]()
		return &empty
	}
	optional, err := With[T](&stream.ts[0])
	if err != nil {
		panic(err)
	}
	return optional
}

func mergeSort[T any](es []T, compare func(a, b T) bool) []T {
	if len(es) < 2 {
		return es
	}
	mid := len(es) / 2
	left := mergeSort(es[:mid], compare)
	right := mergeSort(es[mid:], compare)

	var sorted []T
	i, j := 0, 0
	for len(sorted) != len(es) {
		if i == len(left) {
			sorted = append(sorted, right[j])
			j += 1
		} else if j == len(right) {
			sorted = append(sorted, left[i])
			i += 1
		} else {
			lv := left[i]
			rv := right[j]
			if compare(lv, rv) {
				sorted = append(sorted, lv)
				i += 1
			} else {
				sorted = append(sorted, rv)
				j += 1
			}
		}
	}
	return sorted
}

func (stream Stream[T]) Slice() []T {
	return stream.ts
}
