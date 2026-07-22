package main

import (
	"fmt"
)

func countReports(numSentCh chan int) int {
	total := 0
	for {
		num, ok := <-numSentCh
		if !ok {
			return total
		}
		total += num
	}
}

func main() {
	numSentCh := make(chan int)
	go sendReports(3, numSentCh)
	fmt.Println(countReports(numSentCh))
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}




// func Test(t *testing.T) {
// 	type testCase struct {
// 		numBatches int
// 		expected   int
// 	}

// 	runCases := []testCase{
// 		{3, 114},
// 		{4, 198},
// 	}

// 	submitCases := append(runCases, []testCase{
// 		{0, 0},
// 		{1, 15},
// 		{6, 435},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}

// 	skipped := len(submitCases) - len(testCases)

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		numSentCh := make(chan int)
// 		go sendReports(test.numBatches, numSentCh)
// 		output := countReports(numSentCh)
// 		if output != test.expected {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Test Failed:
//   numBatches: %v
//   expected:   %v
//   actual:     %v
// `, test.numBatches, test.expected, output)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Test Passed:
//   numBatches: %v
//   expected:   %v
//   actual:     %v
// `, test.numBatches, test.expected, output)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// withSubmit is set at compile time depending
// on which button is used to run the tests


// At Textio we're all about keeping track of what our systems are up to with great logging and telemetry.

// The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel. It closes the channel when it's done.

// Complete the countReports function. It should:

// Use an infinite for loop to read from the channel:
// If the channel is closed, break out of the loop
// Otherwise, keep a running total of the number of reports sent
// Return the total number of reports sent
