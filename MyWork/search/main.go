package main

import (
	"fmt"
	"hello/search/binary"
)

func main() {
	searchArr := make([]int, 20)
	for i := 0; i < 20; i++ {
		searchArr[i] = 2 * i
	}

	fmt.Println("Before Search => ", searchArr)

	numIndex := binary.SearchFor(22, 0, len(searchArr), searchArr)

	fmt.Println("numIndex => ", numIndex)
}
