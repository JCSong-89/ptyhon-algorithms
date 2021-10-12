package main

import "fmt"

/*첫 번째 수를 뺀 결과로 키 조회 하면 됨 */

var numsMap = make(map[int]int)

func main(){
	target := 9
	nums := []int{2, 7, 11, 15}

	//키ㅣ와 값을 바꿔 딕셔너리로 저장
	for i, num := range nums{
		numsMap[num] = i
	}

// 타켓에서 첫 번째 수를 뺸 결과로 키 조회 ex: 2 - 9 => 7 
	for i, num := range nums{
		if j, ok := numsMap[target - num]; ok && i != j {
			fmt.Print([]int{i, j})
		}
	}

	fmt.Print([]int{})
}