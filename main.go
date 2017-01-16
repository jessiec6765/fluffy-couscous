package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var output =strings.Fields(s)
	var returnValue = make( map[string]int)


	for _, value := range output {
		if _, found :=returnValue[value];found{
			returnValue[value]++
		}else{
			returnValue[value]=1
		}
	}	
	return returnValue
}

func main() {
	wc.Test(WordCount)
}

