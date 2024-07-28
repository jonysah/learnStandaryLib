package main

import "fmt"

func MapKay[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for key, _ := range m {
		r = append(r, key)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

type list[T any] struct {
	head, tail *element[T]
}

func (t *list[T]) push(val T) {
	if t.tail == nil {
		t.head = &element[T]{val: val}
		t.tail = t.head
	} else {
		t.tail.next = &element[T]{val: val}
		t.tail = t.tail.next
	}
}

func (t *list[T]) getAll() []T {
	r := make([]T, 0, 10)
	p := t.head
	for p != nil {
		r = append(r, p.val)
		p = p.next
	}
	return r
}

func main() {
	m := map[int]string{1: "nih", 2: "你好", 3: "莎娃堤卡"}
	s := MapKay(m)
	fmt.Println(s)
	ls := &list[int]{}
	ls.push(3)
	ls.push(5)
	ls.push(7)
	ls.push(8)
	ls.push(10)
	l := ls.getAll()
	fmt.Println(l)

}
