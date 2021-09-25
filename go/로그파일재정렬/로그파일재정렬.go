package main

import (
	"fmt"
	"sort"
	"strings"
)

/* 로그의 가장 앞 부분은 식별자
문자로 구성된 로그가 숫자 로그보다 앞에 온다.
식별자는 순서 영향에 끼치지 않는다. 문자가 동일할 경우 식별자 순으로 한다.
숫자 로그는 입력 순서대로 한다. */

func dataCleansing (testCase []string) []string {
  cleanLogs := make([]string, 0)

	for i := 0 ; i < len(testCase); i++ {
		sliceStr := strings.Split(testCase[i], " ")
		completStr := sliceStr[1:]
		cleanLogs = append(cleanLogs, completStr[0])
	}
		return cleanLogs
} 

func changeStringInt(str string) string{
	var result = "string"
	slice := strings.Split(str, " ")

	switch slice[0] {
 		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			result = "integer"
	 default:
		result = "string"
 }

	return result
}

func CheckString (testCase []string, cleanLogs []string) []string{
	var strLogs = make([]string, 0)
	var intLogs = make([]string, 0)

	for i := 0 ; i < len(cleanLogs); i++ {

		if changeStringInt(cleanLogs[i]) == "string" {
			strLogs = append(strLogs, testCase[i])
		} else {
			intLogs = append(intLogs, testCase[i])
		}	
	}

	sort.Strings(strLogs)	
	return append(strLogs, intLogs...)
}


func main(){
	var testCase = []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	//result := CheckString(testCase, dataCleansing(testCase))

	//fmt.Print(result)
	fmt.Printf("%#v\n", testCase)
}

