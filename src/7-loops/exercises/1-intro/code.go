package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	const costIncrease float64 = 0.01
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost = totalCost + 1 + costIncrease*float64(i)
	}
	return totalCost
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}
