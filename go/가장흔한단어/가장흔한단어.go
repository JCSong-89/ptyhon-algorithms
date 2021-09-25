package main

import (
	"fmt"
	"strings"
	"container/list"
)

func main() {
	var testCase = "Bob hit a ball, the hit BALL flew far after it was hit."
	var banned = []string{"hit"} 

	lowerCase := strings.Replace(testCase, ".", "", -1)
	lowerCase = strings.Replace(lowerCase, ",", "", -1)
	lowerCase = strings.ToLower(lowerCase)

	var sliceStrPtr *string
	sliceStrPtr = &lowerCase

	banWordInStr(banned, sliceStrPtr)

	sliceStr := strings.Split(lowerCase, " ")
	result := wordCounter(sliceStr, lowerCase)

	fmt.Print(result)
}

func banWordInStr(banned []string, str *string){
	for _, v := range banned{
		*str = strings.Replace(*str, v, "", -1)
	}
}

func wordCounter(words []string, str string) string {
	wordMaps := make(map[string]int)
	queue := list.New()
	queue.PushFront(0)
	var bestWord string

	for _, v := range words{
			wordMaps[v] += 1 
	}

	for key, value := range wordMaps{
		 q := queue.Front()

		if q.Value.(int) == 0 && key != "" {
			queue.PushFront(value)
			bestWord = key
		} else if q.Value.(int) < value && key != "" {
			queue.PushFront(value)
			bestWord = key
		}
	}

	return bestWord
}
