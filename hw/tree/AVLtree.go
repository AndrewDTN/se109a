package main

import (
	"fmt"
)

type AVLtree struct {
	Root *AVLtreeNode
}

type AVLtreeNode struct {
	Value int64
	Times int64
	Height int64
	Left *AVLtreeNode
	Right *AVLtreeNode
}

func NewAVLtree() *AVLtree{
	return new(AVLtree)
}

func (node *AVLtreeNode) Update() {
	if node == nil {
		return
	}

	var leftheight,rightheight int64 = 0,0
	if node.Left != nil {
		leftheight = node.Left.Height
	}
	if node.Right != nil {
		rightheight = node.Right.Height
	}

	maxheight := leftheight
	if rightheight > maxheight {
		maxheight = rightheight
	}
	node.Height = maxheight + 1
}

func (node *AVLtreeNode) Balance() int64 {
	var leftheight,rightheight int64 = 0,0
	if node.Left != nil {
		leftheight = node.Left.Height
	}
	if node.Right != nil {
		rightheight = node.Right.Height
	}
	return leftheight - rightheight
}

func Rightrotation(Root *AVLtreeNode) *AVLtreeNode{
	tem := Root.Left
	tem2 := tem.Right
	tem.Right = Root
	Root.Left = tem2

	Root.Update()
	tem.Update()
	return tem
}

func Leftrotation(Root *AVLtreeNode) *AVLtreeNode {
	tem := Root.Right
	tem2 := tem.Left
	tem.Left = Root
	Root.Right = tem2

	Root.Update()
	tem.Update()
	return tem
}

func LRrotation(node *AVLtreeNode) *AVLtreeNode {
	node.Left = Leftrotation(node.Left)
	return Rightrotation(node)
}

func RLrotation(node *AVLtreeNode) *AVLtreeNode {
	node.Right = Rightrotation(node.Right)
	return Leftrotation(node)
}

//++
func (tree *AVLtree) Add(value int64){
	tree.Root = tree.Root.Add(value)
}

func (node *AVLtreeNode) Add(value int64) *AVLtreeNode{
	if node == nil {
		return &AVLtreeNode{Value: value,Height: 1}
	}
	if node.Value == value{
		node.Times = node.Times + 1
		return node
	}

	var newTreeNode *AVLtreeNode

	if value > node.Value {
		node.Right = node.Right.Add(value)
		factor := node.Balance()
		if factor == -2{
			if value > node.Right.Value {
				newTreeNode = Leftrotation(node)
			} else {
				newTreeNode = RLrotation(node)
			}
		}
	} else {
		node.Left = node.Left.Add(value)
		factor := node.Balance()
		if factor == 2{
			if value < node.Left.Value{
				newTreeNode = Rightrotation(node)
			} else {
				newTreeNode = LRrotation(node)
			}
		}
	}

	if newTreeNode == nil {
		node.Update()
		return node
	} else {
		newTreeNode.Update()
		return newTreeNode
	}
}

//find
func (tree *AVLtree) Findmin() *AVLtreeNode {
	if tree.Root == nil{
		return nil
	}

	return tree.Root.Findmin()
}

func (node *AVLtreeNode) Findmin() *AVLtreeNode{
	if node.Left == nil{
		return node
	}
	return node.Left.Findmin()
}

func (tree *AVLtree) Findmax() *AVLtreeNode {
	if tree.Root == nil{
		return nil
	}
	return tree.Root.Findmax()
}

func (node *AVLtreeNode) Findmax() *AVLtreeNode {
	if node.Right == nil{
		return node
	}
	return node.Right.Findmax()
}

func (tree *AVLtree) Find(value int64) *AVLtreeNode{
	if tree.Root == nil{
		return nil
	}
	return tree.Root.Find(value)
}

func (node *AVLtreeNode) Find(value int64) *AVLtreeNode{
	if value == node.Value {
		return node
	} else if value < node.Value {
		if node.Left == nil{
			return nil
		}
		return node.Left.Find(value)
	} else {
		if node.Right == nil{
			return nil
		}
		return node.Right.Find(value)
	}
}

func (tree *AVLtree) Midorder(){
	tree.Root.Midorder()
}

func (node *AVLtreeNode) Midorder(){
	if node == nil{
		return 
	}

	node.Left.Midorder()

	for i := 0;i <= int(node.Times);i++{
		fmt.Println(node.Value)
	} 

	node.Right.Midorder()
}

//--
func (tree *AVLtree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	tree.Root = tree.Root.Delete(value)
}

func (node *AVLtreeNode) Delete(value int64) *AVLtreeNode{
	if node == nil{
		return nil
	}
	if value < node.Value{
		node.Left = node.Left.Delete(value)
		node.Left.Update()
	} else if value > node.Value{
		node.Right = node.Right.Delete(value)
		node.Right.Update()
	} else {
		if node.Left == nil && node.Right == nil{
			return nil
		}

		if node.Left != nil && node.Right != nil{
			if node.Left.Height > node.Right.Height{
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}

				node.Value = maxNode.Value
				node.Times = maxNode.Times

				node.Left = node.Left.Delete(maxNode.Value)
				node.Left.Update()
			} else {
				minNode := node.Right
				for minNode.Left != nil {
					minNode = minNode.Left
				}
				node.Value = minNode.Value
				node.Times = minNode.Times

				node.Right = node.Right.Delete(minNode.Value)
				node.Right.Update()
			}
		} else {
			if node.Left != nil {
				node.Value = node.Left.Value
				node.Times = node.Left.Times
				node.Height = 1
				node.Left = nil
			} else if node.Right != nil {
				node.Value = node.Right.Value
				node.Times = node.Right.Times
				node.Height = 1
				node.Right = nil
			}
			return node
		}
	}
	var newNode *AVLtreeNode
	if node.Balance() == 2{
		if node.Left.Balance() >= 0{
			newNode = Rightrotation(node)
		} else {
			newNode = LRrotation(node)
		}
	} else if node.Balance() == -2{
		if node.Right.Balance() <= 0{
			newNode = Leftrotation(node)
		} else {
			newNode = RLrotation(node)
		}
	}

	if newNode == nil {
		node.Update()
		return node
	} else {
		newNode.Update()
		return newNode
	}
}

func main() {
	values := []int64{2,3,7,10,10,10,10,23,9,102,109,111,112,113}

	tree := NewAVLtree()
	for _,v := range values{
		tree.Add(v)
	}

	fmt.Println("find min value: ",tree.Findmin())
	fmt.Println("find max value: ",tree.Findmax()) 

	fmt.Println("-------")
	tree.Midorder()
	fmt.Println("-------")
	node := tree.Find(99)
	if node != nil {
		fmt.Println("find 99!")
	} else {
		fmt.Println("not find")
	}

	node = tree.Find(9)
	if node != nil{
		fmt.Println("find 9")
	} else {
		fmt.Println("not find")
	}
	
	tree.Delete(9)
	tree.Delete(10)
	tree.Delete(2)
	tree.Delete(3)
	tree.Add(4)
	tree.Add(3)
	tree.Add(10)
	tree.Delete(111)
	
	node = tree.Find(9)
	if node != nil{
		fmt.Println("find 9")
	} else {
		fmt.Println("not find")
	}
}