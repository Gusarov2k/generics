package generic

import (
	"fmt"
	"strconv"
)


func GenerateString(numbers ...int) string {
	var number string
	for _, val := range numbers {
		number += strconv.Itoa(int(val)) + ", "
	}
	fullLen := len(number)

	return number[:fullLen - 2]
}

func GenerateString32(numbers ...int32) string {
	var number string
	for _, val := range numbers {
		number += strconv.Itoa(int(val)) + ", "
	}
	fullLen := len(number)
	return number[:fullLen - 2]
}

func GenerateStringInter(any ...interface{}) string {

	var number string
	for _, val := range any {
		str := fmt.Sprintf("%v", val)
		number += str + ", "

	}
	fullLen := len(number)
	return number[:fullLen - 2]
}
