package main

import (
	"fmt"
	"strings"
)

// https://qiita.com/drken/items/fd4e5e3630d0f5859067#%E7%AC%AC-9-%E5%95%8F--abc-049-c---daydream-300-%E7%82%B9
// 【キーポイント】
// Greedy アルゴリズム
// 端から決まって行く Greedy アルゴリズム
// 後ろから解く
func main() {
	// https://atcoder.jp/contests/abs/tasks/arc065_a
	var s string
	fmt.Scan(&s)
	s = strings.Replace(s, "dream", "D", -1)
	s = strings.Replace(s, "erase", "E", -1)
	s = strings.Replace(s, "Der", "", -1)
	s = strings.Replace(s, "Er", "", -1)
	s = strings.Replace(s, "D", "", -1)
	s = strings.Replace(s, "E", "", -1)
	s = strings.TrimSpace(s)

	if s == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
