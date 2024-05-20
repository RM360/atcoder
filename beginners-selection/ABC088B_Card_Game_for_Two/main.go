package main

import (
	"fmt"
	"sort"
)

func main() {

	var N int
	fmt.Scanf("%d", &N)

	cards := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &cards[i])
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[j] < cards[i]
	})

	var alice int
	var bob int

	for i := range cards {

		if i == 0 {
			alice += cards[i]
		} else {
			if i%2 == 1 {
				bob += cards[i]
			} else {
				alice += cards[i]
			}
		}
	}

	fmt.Println(alice - bob)
}
