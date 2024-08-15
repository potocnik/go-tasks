package utils

import "fmt"

func Check(err error) {
	if err != nil {
		fmt.Println("Caugth managed error", err)
		panic(err)
	}
}

func CheckWithMessage(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		panic(err)
	}
}
