package main

import (
	"fmt"
	"sort"
)

/* 배열을 입력받아 합으로 0을 만들 수 있는 3개의 엘리먼트를 출력하라 */
/* 투 포인터로 합 계산 */
/*중복값 제거해주면서 앞과 끝에서  */
/* 합이 0보다 작으면 left 우측으로 이동, 0보다 크다면 right로*/

var nums = []int{-1, 0, 1, 2, -1, -4}

func main() {
	results := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums) -2; i++ {
		if i > 0 && nums[i] == nums[i - 1]{
			continue
		} // 중복된 값 건너뛰기

		left, right := i + 1, len(nums) - 1
		for; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left += 1
			} else if sum > 0{
				right -= 1
			} else {
				results = append(results, []int{nums[i], nums[left],nums[right]})

				for; right != len(nums) - 1 &&left < right && nums[left] == nums[right + 1]; {
					left += 1
				}
				for; left < right && nums[right] == nums[right - 1]; {
					right -= 1
		  	}
				left += 1
				right -= 1
		}
	}
 }
 fmt.Print(results)
}