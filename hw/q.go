package main

import (
    "fmt"
)

type IntNode struct {
    value    int
    //previous *IntNode
    next     *IntNode
}

type MyList struct {
    head *IntNode
    tail *IntNode
}

func (q *MyList) Enqueue(n int) {
    new := IntNode{value: n}
    if q.head == nil {
        q.head = &new
        q.tail = &new
    } else {
        q.tail.next = &new
        q.tail = &new
    }
}

func (q *MyList) Dequeue() int {
    result := q.head
    q.head = q.head.next
    return result.value
}

func main() {
    var q MyList
    q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
    fmt.Printf("FirstInFirstOut head: %v at memory address %p \n", q.head.value, q.head)
    fmt.Printf("FirstInFirstOut tail: %v at memory address %p \n", q.tail.value, q.tail)
    fmt.Println("dequeuing a value: ", q.Dequeue())
    fmt.Printf("FirstInFirstOut head: %v at memory address %p \n", q.head.value, q.head)
    fmt.Printf("FirstInFirstOut tail: %v at memory address %p \n", q.tail.value, q.tail)
}