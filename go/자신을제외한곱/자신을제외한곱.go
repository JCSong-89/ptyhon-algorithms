package main

import (
	"container/list"
	"fmt"
)

var sum int = 1
var out []int

/* 나를 제외한 나머지 모든 요소의 곱셈 결과 */
func main() {
	test := []int{1, 2, 3, 4}
	out = []int{}
	q := list.New()
	q.PushFront(sum)

	sums(0, test, q)

	fmt.Print(out)
}

func sums(n int, t []int, q *list.List) {
	for i := 0; i < len(t); i++ {
		if n == i {
			continue
		}
			v := q.Front().Value.(int)
			sum = t[i] * v 

		fmt.Println(sum)
		q.PushFront(sum)
	}

	out = append(out, sum) 
	sum = 1
	q.PushFront(sum)

	if n == len(t) -1 {
		return 
	}

	sums(n + 1, t, q)
}
