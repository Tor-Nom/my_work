package opt

import "fmt"

type StringReverse struct {
	OptString string
}

func NewStringReverse(optString string) *StringReverse {

	fmt.Println("Before Reverse => ", optString)
	return &StringReverse{
		OptString: optString,
	}
}

func (sr *StringReverse) Reverse() {
	stringSlice := []byte(sr.OptString)
	length := len(stringSlice)
	for i := 0; i < length/2; i++ {
		stringSlice[i], stringSlice[length-1-i] = stringSlice[length-1-i], stringSlice[i]
	}
	fmt.Println("After StringSlice => ", string(stringSlice))

}
