package utils

import "fmt"

func HandleError(message string, errorFunction string) error {
	fmt.Println(message + " " + errorFunction)
	return fmt.Errorf(message + " " + errorFunction)
}
