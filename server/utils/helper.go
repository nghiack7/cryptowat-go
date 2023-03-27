package utils

import (
	"fmt"
	"strconv"
)

func ParseFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return f
	}
	return f
}
