package main
import (
	"fmt"
)

type Node struct{
	data int
	next *Node
}

type List struct{
	head *Node
	tail *Node
}

func (q *List) Enqueue(n int){
	new := &Node{n,nil}
	if q.head == nil{
		q.head = new
		q.tail = new
	} else {
		q.tail.next = new
		q.tail = new
	} 
}

func (q *List) Dequeue() int{
	out := q.head.data
	q.head = q.head.next
	return out
}

func (q *List) Show() {
	temp := q.head
	for temp != nil {
	  fmt.Println(temp.data)
	  temp = temp.next
	}
}

func main(){
	var q List
	q.Enqueue(2)
	q.Show()
	fmt.Println("------")
	q.Enqueue(5)
	q.Show()
	fmt.Println("------")
	q.Enqueue(6)
	q.Show()
	fmt.Println("------")
	q.Enqueue(4)
	q.Show()
	fmt.Println("------")
	q.Dequeue()
	q.Show()
	fmt.Println("------")
	q.Dequeue()
	q.Show()
}