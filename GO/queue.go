// 在 Go 中实现一个基本的队列（Queue）可以使用切片（slice）来存储数据。支持常见的队列操作，如 Enqueue（入队），Dequeue（出队）和 Peek（查看队列头部元素）。
package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Equeue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Denqueue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("is empty")
	}
	element := q.items[0]
	q.items = q.items[1:]
	fmt.Println(q.items)
	return element, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("is empty")
	}
	return q.items[0], nil
}

func main() {
	q := &Queue{}
	q.Equeue(10)
	q.Equeue(20)
	q.Equeue(30)

	peek, _ := q.Peek()
	fmt.Println(peek)

	dep, _ := q.Denqueue()
	fmt.Println(dep)

	dep, _ = q.Denqueue()
	fmt.Println(dep)

}
