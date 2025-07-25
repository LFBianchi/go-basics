package main

import (
	"fmt"
)

func maxMessages(thresh float64) int {
	const costIncrease float64 = 0.01
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost = totalCost + 1 + costIncrease*float64(i)
		if totalCost > thresh {
			return i
		}
	}
}

// don't edit below this line

func test(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(40.00)
	test(50.00)
}
