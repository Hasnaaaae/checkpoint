package main

import (
	"fmt"
)

func CountAlpha(s string) int {
compt := 0
for _ , ss := range s {
	if (ss >= 65 && ss <= 90) || (ss >= 97 && ss <= 122) {
        compt++
	}
}
return compt
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
