package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("defer one")
	}()
	funPanic()
}

func funPanic() {
	//fmt.Println("defer one")
	panic("panic one")
}
