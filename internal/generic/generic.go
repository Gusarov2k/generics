package generic

import (
	"log"
	"reflect"
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

	fn := reflect.ValueOf(any)
	fnType := fn.Type()

	log.Println(fnType)
	//var number string
	//if fnType == int {
	//	for _, val := range any {
	//		number += strconv.Itoa(int(val)) + ", "
	//	}
	//	fullLen := len(number)
	//
	//	return number[:fullLen - 2]
	//
	//}
	//log.Println(any.)
	//var number string
	//for _, val := range any {
	//	number += strconv.Itoa(int(val)) + ", "
	//}
	//fullLen := len(number)
	//return number[:fullLen - 2]
	return "ze"
}
