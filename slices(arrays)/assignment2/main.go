package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	if plan==planPro{
		return messages[0:3],nil
	}else if plan==planFree{
		return messages[0:2],nil
	}else{
		return messages[0:0],errors.New("unsupported plan")
	}
}
func main(){
	msg:=[3]string{"hello","bye","show msg"}
	fmt.Println(getMessageWithRetriesForPlan("free",msg))
	fmt.Println(getMessageWithRetriesForPlan("pro",msg))
	fmt.Println(getMessageWithRetriesForPlan("",msg))
}
// Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input as well as an array of 3 messages. You've been provided with constants representing the plan types at the top of the file.

// If the plan is a pro plan, return all the strings from the messages input in a slice.
// If the plan is a free plan, return the first 2 strings from the messages input in a slice.
// If the plan isn't either of those, return a nil slice and an error that says unsupported plan.