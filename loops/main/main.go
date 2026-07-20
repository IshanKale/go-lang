package main

import "fmt"
func bulkSend(numMessages int) float64 {
	cost := 0.00
	ans := 0.0
	for i := 0; i < numMessages; i++ {
		ans += cost
		cost += 0.01
	}
	return ans
}

func main() {
	fmt.Println("bulkSend(0) =", bulkSend(0))
	fmt.Println("bulkSend(1) =", bulkSend(1))
	fmt.Println("bulkSend(5) =", bulkSend(5))
}

