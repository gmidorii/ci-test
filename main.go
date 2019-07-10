package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(a int) string {
	if (a % 15) == 0 {
		return "fizzbuzz"
	}
	if (a % 5) == 0 {
		return "buzz"
	}
	if (a % 3) == 0 {
		return "fizz"
	}
	return strconv.Itoa(a)
}

func main() {
	for i := 1; i < 30; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
