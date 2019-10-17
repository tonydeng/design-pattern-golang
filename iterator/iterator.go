package iterator

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	First()
	isDone() bool
	Next() interface{}
}

type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumberIterator{
		numbers: n,
		next:    n.start,
	}
}

type NumberIterator struct {
	numbers *Numbers
	next    int
}

func (i *NumberIterator) First() {
	i.next = i.numbers.start
}

func (i *NumberIterator) isDone() bool {
	return i.next > i.numbers.end
}

func (i *NumberIterator) Next() interface{} {
	if !i.isDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.isDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}
