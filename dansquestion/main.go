package main

import "fmt"

type dan struct {
	num1 int
	num2 int
	num3 int
	num4 int
	num5 int
}

func (d dan) getSum() int {
	return d.num1 + d.num2 + d.num3 + d.num4 + d.num5
}

func main() {
	dansStruct := dan{1, 2, 3, 4, 5}

	fmt.Println(dansStruct.getSum())
}
