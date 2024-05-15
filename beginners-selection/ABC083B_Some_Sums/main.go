package main

import (
	"fmt"
	"strconv"
)

func main() {

	var N, A, B int
	fmt.Scanf("%d %d %d", &N, &A, &B)

	total := 0

	for i := 1; i <= N; i++ {
		str_i := strconv.Itoa(i) // 整数を文字列に変換
		sum := 0
		for _, char := range str_i {
			//各桁の総和sumを求める
			digit, err := strconv.Atoi(string(char))

			if err != nil {
				return
			}

			sum += digit
		}

		if A <= sum && sum <= B {
			total += i
		}
	}

	fmt.Println(total)
}
