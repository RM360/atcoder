package main

import (
	"fmt"
	"sort"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	mochi := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &mochi[i])
	}

	sort.Slice(mochi, func(i, j int) bool {
		return mochi[j] < mochi[i]
	})

	count := 0
	var shitano_mochi int
	for i := range mochi {

		if i == 0 {
			count += 1
		} else {
			if shitano_mochi > mochi[i] {
				count += 1
			}
		}
		shitano_mochi = mochi[i]
	}

	fmt.Println(count)
}
