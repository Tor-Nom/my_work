package quick

import "fmt"

type MyQuick struct {
	SortArr []int
	Length  int
}

func NewMyQuick(sortArr []int) (*MyQuick, bool) {
	fmt.Println("Before MyQuick => ", sortArr)
	arrLen := len(sortArr)
	if arrLen == 0 || arrLen == 1 {
		return nil, false
	}
	return &MyQuick{
		SortArr: sortArr,
		Length:  arrLen,
	}, true
}

func (mq *MyQuick) Sort() {
	QuickSort(mq.SortArr)
}

func (mq *MyQuick) Print() {
	fmt.Println("After QuickSort => ", mq.SortArr)
}

func NoFor() {

}

func QuickSort(sortArr []int) {
	left, right := 0, len(sortArr)-1
	if left >= right {
		return
	}

	flag, flagIndex := sortArr[0], 0

	for left < right {
		for sortArr[right] >= flag && right > left {
			right--
		}

		if right > left {
			sortArr[flagIndex] = sortArr[right]
			flagIndex = right
		}

		for sortArr[left] <= flag && right > left {
			left++
		}

		if right > left {
			sortArr[flagIndex] = sortArr[left]
			flagIndex = left
		}

		//for right > left {
		//	if sortArr[right] > flag {
		//		right--
		//	} else {
		//		sortArr[flagIndex] = sortArr[right]
		//		flagIndex = right
		//		break
		//	}
		//}
		//for right > left {
		//	if sortArr[left] < flag {
		//		left++
		//	} else {
		//		sortArr[flagIndex] = sortArr[left]
		//		flagIndex = left
		//		break
		//	}
		//}
	}
	sortArr[flagIndex] = flag
	QuickSort(sortArr[:flagIndex])
	QuickSort(sortArr[flagIndex+1:])
	return
}
