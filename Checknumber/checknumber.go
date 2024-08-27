package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
for _, ar := range arg {
	if ar >= 48 && ar <= 57 {
		return true
	}
}
return false	
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
