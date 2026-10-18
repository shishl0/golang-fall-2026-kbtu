package main

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Push(v T) *List[T] {
	node := &List[T]{val: v}
	if l == nil {
		return node
	}
	cur := l
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
	return l
}

func (l *List[T]) Len() int {
	count := 0
	cur := l
	for cur != nil {
		count++
		cur = cur.next
	}
	return count
}

func (l *List[T]) String() string {
	var parts []string
	cur := l
	for cur != nil {
		parts = append(parts, fmt.Sprint(cur.val))
		cur = cur.next
	}
	return "[" + strings.Join(parts, " -> ") + "]"
}

func task7() {
	var l *List[int]
	fmt.Println(l, "len:", l.Len())

	l = l.Push(1)
	l = l.Push(2)
	l = l.Push(3)
	fmt.Println(l, "len:", l.Len())

	var words *List[string]
	words = words.Push("go")
	words = words.Push("is")
	words = words.Push("fun")
	fmt.Println(words, "len:", words.Len())
}
