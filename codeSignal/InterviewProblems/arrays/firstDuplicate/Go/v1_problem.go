package main

import (
	"fmt"
	"math"
)

func firstDuplicate(a []int)int{

	//Account for <=
	dupArray := make([]int, len(a)+1)
	for x := range dupArray{
		dupArray[x] = -1
	}

	smallestValue := int(math.Pow(10, 5))

	for x := range a{
		if -1 == dupArray[a[x]]{
			dupArray[a[x]] = 0
		} else if 0 == dupArray[a[x]]{
			dupArray[a[x]] = x
			if x < smallestValue{
				smallestValue = x
			}
		}
	}

	if smallestValue == int(math.Pow(10, 5)) {
		smallestValue = -1
	} else {
		smallestValue = a[smallestValue]
	}

	return smallestValue
}

func main(){

	myNumbers := []int{2, 1, 3, 5, 3, 2}

	myResult := firstDuplicate(myNumbers)

	fmt.Println(myResult)
}
