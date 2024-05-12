package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d", &n)

    numbers := make([]int, n)
	for i := 0; i < n; i++ {
        fmt.Scanf("%d", &numbers[i])
    }

    count := 0
    
    for {
        if(containsOdd(numbers)){
            fmt.Printf("%d", count)
            break
        }
        
        for i := range numbers{
            numbers[i] /= 2
        }

        count ++
    }
}

func containsOdd(numbers[] int) bool {
    for _, num := range numbers {
        if num%2 != 0 {
            return true
        }
    }
	return false
}