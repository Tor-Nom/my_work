package main

import "hello/stringOpt/opt"

func main() {
	sr := opt.NewStringReverse("abcdefgh")

	sr.Reverse()
}
