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

func count(ch chan int) {
	fmt.Println("ch = > ", ch)
	ch <- 1
}
