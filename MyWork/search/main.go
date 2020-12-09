package main

import (
	"fmt"
	"hello/search/binary"
	"hello/search/search_int"
)

func main() {
	searchArr := make([]int, 20)
	for i := 0; i < 20; i++ {
		searchArr[i] = 2 * i
	}

	fmt.Println("Before Search => ", searchArr)

	numIndex := binary.SearchFor(22, 0, len(searchArr), searchArr)

	fmt.Println("numIndex => ", numIndex)

	optString := "ee1039ccc254ad222321qaads22321341wwwq1039a"

	intCount := search_int.GetIntCount(optString)

	fmt.Println("intCount => ",intCount)
}
