package main

import (
	"container/list"
	"fmt"
	"reflect"
)

/* 연결 리스트가 팰린드롬 구조인지 판별 */

type Node struct {
	next *Node
	data int
}

func (n *Node) NextNode() *Node{
	return n.next
}

var q *list.List = list.New()
var start []int 
var end []int

func main() {
	test := []int{1, 2}
	nodes := []*Node{}
	start = []int{}
	end = []int{}

	first := Node{data: test[0]}
	nodes = append(nodes, &first)
	q.PushFront(&first)

	for i :=1; i < len(test); i++  {
		prevNode := q.Front().Value.(*Node)
		node := Node{data: test[i]}
		prevNode.next = &node
		nodes = append(nodes, &node)
		q.PushFront(&node)
	}

	searchNode(&first)
	if len(start) == 2 {
		fmt.Println(false)
	} else {
		fmt.Println(reflect.DeepEqual(start, end))
	}
}


func searchNode (node *Node) {
	start = append(start, node.data)
	nextNode := node.NextNode()
	if nextNode != nil {
			searchNode(nextNode)
	}

	end = append(end, node.data)	
}