package utils

import "fmt"

func Logger(value interface{}) string {
	s := fmt.Sprintf("It is %v", value)
	return s
}
