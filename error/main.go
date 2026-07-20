package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	cost1, err1 := sendSMS(msgToCustomer)
	if err1 != nil {
		return 0, err1
	}
	cost2, err2 := sendSMS(msgToSpouse)
	if err2 != nil {
		return 0, err2
	}
	return cost1 + cost2, nil

}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func main() {
	msgToCustomer := "hello"
	msgToSpouse := "fljwoighioewhwhtpwhtpihrwiopghprwghprwhgprwhgrwhghrwp9gh-9rwhg"
	res, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)

	}
}

// We offer a product that allows businesses to send pairs of messages to couples. It is mostly used by flower shops and movie theaters.

// Complete the sendSMSToCouple function. It should send 2 messages, first to the customer and then to the customer's spouse.

// Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0 and the error.
// Do the same for the msgToSpouse
// If both messages are sent successfully, return the total cost of the messages added together.
