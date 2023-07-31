package utils

import "fmt"

func LogErrors(errors ...error) {
	if len(errors) == 0 {
		return
	}

	for _, err := range errors {
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}
