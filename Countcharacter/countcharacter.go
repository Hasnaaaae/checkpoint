package main

import (
	"fmt"
)

func CountChar(str string, c rune) int {
	var compt int
	for   _ , s := range str {
		if s == c {
			compt++
		}
	}
   return compt
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}