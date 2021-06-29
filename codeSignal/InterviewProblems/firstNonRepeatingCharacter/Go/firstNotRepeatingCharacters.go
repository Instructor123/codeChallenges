package main

import (
	"strings"
)

func firstNotDuplicateCharacter(s string) string{

	var temp string

	for x := range s{
		if 1 == strings.Count(s, string(s[x])){
			temp = string(s[x])
			break
		}
	}

	if 0 == len(temp) {
		temp = "_"
	}

	return temp
}

func main(){
	myString := string("abacabaabacaba")

	firstNotDuplicateCharacter(myString)
}
