package main

import (
	"fmt"
)

// 回答を見た
// https://atcoder.jp/contests/abs/tasks/abc085_c/editorial
func main() {
	var N, Y int
	fmt.Scanf("%d %d", &N, &Y)

	num10000, num5000, num1000 := -1, -1, -1

	for a := 0; a <= N; a++ {
		for b := 0; b+a <= N; b++ {

			c := N - a - b
			total := 10000*a + 5000*b + 1000*c

			if total == Y {
				num10000 = a
				num5000 = b
				num1000 = c
			}
			// 1 秒間で処理できる for 文ループの回数は、10^8回程度
			// cの枚数はa、bが決まればcは決まるので以下のループが不要になる。
			// for c := 0; c <= N; c++ {
			// 	total := 10000*a + 5000*b + 1000*c

			// 	if total == Y {
			// 		num10000 = a
			// 		num5000 = b
			// 		num1000 = c
			// 	}
			// }
		}
	}

	fmt.Println(num10000, num5000, num1000)
}
