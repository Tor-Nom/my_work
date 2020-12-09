package main

import (
	"fmt"
	"log"
	"runtime"
)

type Base struct {
	Name string
	test string
}

//type File interface {
//	Read()
//	Write()
//}

type Rcfile interface {
	Read()
}

type Cfile struct {
}

func (c *Cfile) Read() {

}
func (c *Cfile) Write() {

}

func (b *Base) Fool() {

}

func (b *Base) Bar() {

}

type Fool struct {
	Base
	*log.Logger
	Name string
}

func (f *Fool) Bar() {
	f.Base.Bar()
	f.Logger.Fatal(222)
	f.Name = "2"
	f.Base.Name = "3"
}

func main() {

	GetIntCount()
	return
	s := Search([]string{"(","("})
	fmt.Println("s => ",s)

	defer fmt.Println("333333")

	defer func() {
		defer fmt.Println("222222")
	}()

	defer fmt.Println("11111")
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum =>", cpuNum)
	intChan := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		intChan[i] = make(chan int, 0)

		go count(intChan[i])
	}

	for _, ch := range intChan {
		<-ch
	}
}

type IntCount struct {
	num string
	cnt int
}

func GetIntCount()  {
	optString := "ee1039ccc254ad222321qaads22321341wwwq1039a"
	var intCount []IntCount
	strCnt := len(optString)
	for i := 0;i<strCnt;i++ {
		runeString := rune(optString[i])

		var intStart = i
		var intString string
		if runeString >= 48 && runeString <= 57 {
			for ;intStart <strCnt;intStart++ {
				i = intStart
				runeIntString := rune(optString[intStart])
				if runeIntString >= 48 && runeIntString <= 57 {
					intString = intString + string(runeIntString)
				}else {
					index := InArrayString(intString,intCount)
					if index >= 0{
						intCount[index].cnt++
					}else {
						intCount = append(intCount,IntCount{
							num : intString,
							cnt: 1,
						})
					}
					break
				}
			}

		}
	}
	fmt.Println("intString => ",intCount)

}
func count(ch chan int) {
	fmt.Println("ch = > ", ch)
	ch <- 1
}

func InArrayString(idle string,array []IntCount) int {
	for i ,v := range array {
		if idle == v.num {
			return i
		}
	}
	return -1
}
