package main

import (
	"fmt"
)

func main() {
	i := 2
	defer test(i)
	i = 4
	defer test(i)
	i = 5
	defer test(i)
}

func test(i int) {
	fmt.Println("test i : ", i)
}
