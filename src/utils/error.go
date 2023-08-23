package utils

import "fmt"

func HandleError(message string, errorFunction string) {
	fmt.Println(message + " " + errorFunction)
}
