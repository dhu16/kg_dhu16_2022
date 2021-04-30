package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	phoneticEquivalent := make(map[rune]string)
	phoneticEquivalent = map[rune]string {
		'1' : "One",
		'2' : "Two",
		'3' : "Three",
		'4' : "Four",
		'5' : "Five",
		'6' : "Six",
		'7' : "Seven",
		'8' : "Eight",
		'9' : "Nine",
		'0' : "Zero",
	}
	
	nums := os.Args[1:]
	var res []string
	
	for _, s := range nums {
		var sb strings.Builder
		for _, c := range s {
			str := phoneticEquivalent[c]
			sb.WriteString(str)
		}
		res = append(res, sb.String())	
	}
	
	print(res)
}

func print(arr []string) {
	for i, s := range arr {
		if i < len(arr)-1 {
			fmt.Print(s + ", ")
		} else {
			fmt.Print(s)
		}
	}
}