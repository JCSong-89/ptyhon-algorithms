package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	testCase := []string{"eat", "tea", "tan", "ate", "nat", "bat"}	
	sortWords := make([]string, 0)

	for i :=0; i < len(testCase); i++ {
		slice := strings.Split(testCase[i], "")
		sort.Strings(slice)
		sortWords = append(sortWords, strings.Join(slice, ""))
	}

	wordMaps := make(map[string]int)
	count := 1

	for i := 0; i < len(testCase); i++ {
		if _, ok := wordMaps[sortWords[i]]; !ok {
			wordMaps[sortWords[i]] = count
			count += 1
		} 
	}

	resultMaps := make(map[string]int)

	for i := 0; i < len(testCase); i++ {
		for key, value := range wordMaps{
			if checkAnagram(testCase[i], key) {
				resultMaps[testCase[i]] = value
			}
		}
	}
	
	wordArray := make([][]string, len(wordMaps))

	for key, value := range resultMaps{
		wordArray[value -1] = append(wordArray[value-1], key)
	}

	fmt.Print(wordArray)
}

func checkAnagram(str string, mapKey string) bool{
		slice := strings.Split(str, "")
		sort.Strings(slice)
		joinStr := strings.Join(slice, "")

		if joinStr == mapKey {
			return true
		} 

		return false
	}