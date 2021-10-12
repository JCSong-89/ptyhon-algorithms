package main

import "fmt"

/* 한번 거래로 낼 수 있는 최대 이익 */
/* 최소값, 최대값을 계속 갱신하는 알고리즘 */
var min_price int
var max_price int

func main() {
	test := []int{7, 1, 5, 3, 6, 4}
	money := 0
	max_price = 0
	min_price = test[0]

	for i :=0; i < len(test); i++ {
		if test[i] < min_price {
			min_price = test[i]
		}

		if i != 0 && test[i] > max_price {
			max_price = test[i]
		}

		money = max_price - min_price
	}

	fmt.Println(money)
}