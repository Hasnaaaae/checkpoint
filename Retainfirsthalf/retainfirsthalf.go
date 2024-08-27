package main

import (
	"fmt"
)

func RetainFirstHalf(str string) string {
    var half, word, word2 string 
	if len(str) == 1 || len(str) == 0 {
		return str
	}
    tab := Splate(str)
	if len(tab) % 2 == 0 {
		for i := 0; i<len(tab); i++ {
			half = half + tab[i] + " "
            if i == (len(tab)/2)-1 {
				half = half[0:len(half)-1]
                break
			}
		}
	}else {
		for i := 0; i<len(tab); i++ {
			half = half + tab[i] + " "
            if i == (len(tab)/2)-1 {
				word = tab[i+1]
                break
			}
			
		}
		for  j := 0; j<len(word) ; j++ {
			word2 = word2 + string(word[j]) 
			if j == (len(tab)/2)-1 {
                break
			}
		} 
		half = half + word2

	}
	return half
}

func Splate(s string) []string {
	var tab []string
	var ss string
    for i := 0; i<len(s); i++ {
		if s[i] != ' ' {
            ss = ss + string(s[i])
		}else {
			tab = append(tab, ss)
			ss = ""
		}
		if i == len(s)-1 {
			tab = append(tab, ss)
		}
	}
	return tab 
}


func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}