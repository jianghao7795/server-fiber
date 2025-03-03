package utils

import "fmt"

func Logger(value any) string {
	s := fmt.Sprintf("It is %v", value)
	return s
}
