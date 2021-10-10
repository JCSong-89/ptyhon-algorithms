package main

import (
	//"container/list"
	"fmt"
 "strconv"
)

/*
n개의 팀이 주어짐.
한명의 팀은 다른 팀과의 리그를 펼침

리그에서 치루어지는 경기수

1명 x n개 / 2
*/

func main(){
	ints, _ := strconv.Atoi(ten)
	sum := 0
	n := 20
	sumPtr := &sum

	for i :=1; i <= 10; i++ {
		if n % i == 0 {
			*sumPtr = *sumPtr + i
		}
	}

	s := string(sum)
	fmt.Println(s)
}