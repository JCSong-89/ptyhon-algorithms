package main

import (
	"container/list"
	"fmt"
	"strings"
)
func main() {
	testCase := "cbbd"
	wordSlice := strings.Split(testCase, "")
	var topWord string
	queue :=  list.New()
	queue.PushBack(0)

	for i := 0; i < len(wordSlice); i++ {
		count := strings.Count(testCase, wordSlice[i]) 
		if queue.Front().Value.(int) < count { 
			queue.PushFront(count)
			topWord = wordSlice[i]
		}
	}

	topWordIdx := make([]int, 0)

	for i := 0; i < len(wordSlice); i++ {
		if wordSlice[i] == topWord{
			topWordIdx = append(topWordIdx, i)
		}
	}
	
	fmt.Print(testCase[topWordIdx[0]:topWordIdx[len(topWordIdx)-1]+1])
}