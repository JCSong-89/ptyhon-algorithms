package main

import (
	"fmt"
	"sort"
)

// n개의 페어를 이용한 min(a, b)의 합으로 만들 수 있는 가장 큰 수를 출력
/* 정렬된 상태에서 짝수 번째 값을 더하면 될듯 */
func  main()  {
	nums := []int{2, 3, 1, 4}
	sort.Ints(nums)
	sum := 0
	for i := 0; i <len(nums); i++ {
		if i % 2 == 0 {
			sum += nums[i]
		}
	} 

	fmt.Print(sum)
}