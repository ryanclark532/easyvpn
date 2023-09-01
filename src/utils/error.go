package utils

import "fmt"

func HandleError(err error, errorFunction string) {
	fmt.Println(err.Error() + " " + errorFunction)
}
