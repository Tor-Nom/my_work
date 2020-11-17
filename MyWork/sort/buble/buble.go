package buble

import "fmt"

type MyBuble struct {
	SortArr []int
}

func NewMyBuble(sortArr []int) *MyBuble {
	return &MyBuble{
		SortArr: sortArr,
	}
}

func (mb *MyBuble) Sort() {
	n := len(mb.SortArr)
	if n == 1 || n == 0 {
		return
	}

	var swapFlag bool
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if mb.SortArr[j] > mb.SortArr[j+1] {
				swapFlag = true
				mb.SortArr[j+1], mb.SortArr[j] = mb.SortArr[j], mb.SortArr[j+1]
			}
		}

		fmt.Println("第v%次排序完：", i, mb.SortArr)
		if !swapFlag {
			return
		}
	}
	return
}
