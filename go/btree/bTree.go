package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Tree struct {
	Left * Tree
	Value int
	Right *Tree
} // 트리 구조 

func traverse(t *Tree) {
	if t == nil {
		return 
	}

	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
} // 재귀호출로 모든 노드 방문

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	} // 노드가 비어있을 경우 루트 로드로 한다.

	if v == t.Value {
		return t
	} // 해당 값이 노드 벨류와 같다면 이미 노드가 정리 된것으로 판단

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	} // 작으면 왼쪽

	t.Right = insert(t.Right, v)
	return t // 크면 오른쪽
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	} 

	return t
} // 난수 노드 생성기


func main() {
	tree := create(10)
	traverse(tree)
	fmt.Println(tree.Value)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
}