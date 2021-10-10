package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
type Node interface {
	Next()
	Prev()
}

*/
type Node struct {
	Alpabet string
 Next *Node
 Prev *Node
}

func (n *Node) CheckNode(s string, c *int) int{
	if n.Alpabet != s && n.Next != nil{
		n.Next.CheckNode(s, c)
	}	else if n.Alpabet != s &&  n.Next == nil {
		*c += 1
		return *c 
	}
	return *c 
}

func solution(begin string, target string, words []string) int {
	if begin == "" {
		return 0
	}
	if target == "" {
		return 0
	}

	wordMap := make(map[string]int)

	for _, v := range words {
		wordMap[v] = 1
	}

	if wordMap[target] != 1 {
		return 0
	}

	startWordSlice := strings.Split(begin, "")
	targetWordSlice := strings.Split(target, "")
	sort.Strings(startWordSlice)
	sort.Strings(targetWordSlice)

	nodes := make([]*Node, 0)
	nodesPtr := &nodes

	for i := 1; i < len(targetWordSlice); i++{
		if i == 1 {
			firstNode := Node{Prev: nil, Alpabet: targetWordSlice[i-0]}
			secNode := Node{Prev: &firstNode, Alpabet: targetWordSlice[i]}
			firstNode.Next = &secNode
			*nodesPtr = append(*nodesPtr, &firstNode, &secNode)
		}

		if i != len(target) - 1 {
			node := nodes[i]
			nextNode :=  Node{Prev: node, Alpabet: targetWordSlice[i + 1], Next:(*nodesPtr)[i-1]}
			node.Next = &nextNode
			*nodesPtr = append(*nodesPtr, &nextNode)
		} else {
			node := nodes[i]
			node.Next = nil
		}
	}
	count := 1
	countPtr := &count

	for _, v := range startWordSlice {
			c := 0
		*countPtr += nodes[0].CheckNode(v, &c)
	}

	return count
}

func main(){
	words := []string{"hhh", "dot", "dog", "lot", "log", "hit"}
	begin := "hit"
	target := "lot"
	fmt.Println(solution(begin, target, words))
}