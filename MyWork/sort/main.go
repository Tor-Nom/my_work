package main

import (
	"fmt"
	"hello/sort/insertion"
	"math/rand"
	"sort"
)

type MySort interface {
	Sort()
	Print()
}

func main() {
	mySortArr := make([]int, 20)
	for i := 0; i < 20; i++ {
		mySortArr[i] = rand.Intn(500)
	}

	//mySortArr = []int{10, 8, 22, 9, 7, 20, 31, 13}
	//mySortArr = []int{81, 387, 347, 59, 81, 318, 425, 40, 456}
	//冒泡
	//buble.Sort(sortArr)
	//选择
	//sortStruct, needNew := _select.NewMySelect(mySortArr)
	//插入排序
	sortStruct, needNew := insertion.NewMyInsertion(mySortArr)
	//归并排序
	//sortStruct, needNew := merge.NewMyMerge(mySortArr)
	//sortStruct, needNew := quick.NewMyQuick(mySortArr)
	if needNew {
		sortStruct.SortOpt()
		sortStruct.Print()
	}
}

func SysSort(sortArr []int) {
	sort.Ints(sortArr)
	fmt.Println(sortArr)
}
