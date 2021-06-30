package main

import "fmt"

/*
	When we rotate, we need to know the maximum size of the array because the edges will depend on it.
	(0,0) -> (3,0)	|	(1,0) -> (3,1)	|	(2,0) -> (3,2)	|	(3,0) -> (3,3)
	(0,1) -> (2,0)	|	(1,1) -> (2,1)	|	(2,1) -> (2,2)	|	(3,1) -> (2,3)
	(0,2) -> (1,0)	|	(1,2) -> (1,1)	|	(2,2) -> (1,2)	|	(3,2) -> (1,3)
	(0,3) -> (0,0)	|	(1,3) -> (0,1)	|	(2,3) -> (0,2)	|	(3,3) -> (0,3)
 */
/*
	//rotate left
	for x := 0; x < arraySize; x++{
		for y := arraySize-1; y >= 0; y--{
			tempValue := a[y][x]
			temp[y][x] = a[x][y]
			temp[x][y] = tempValue
			fmt.Println(tempValue)
		}
	}
 */
func rotateImage(a [][]int) [][]int{

	arraySize := len(a)

	for i := 0; i < int(arraySize/2); i++{
		placeHolder := (i+1) % arraySize
		for j := i; 0 != (arraySize - placeHolder) - 1; j++{
			placeHolder = (j+i+1) % arraySize
			local := arraySize - (i % arraySize)

			//0,0
			temp1 := a[i][j]
			//0,3
			temp2 := a[j][local-1]
			//3,3
			modValue := local - (j % local)
			temp3 := 0
			temp4 := 0
			if i != 0 {
				temp3 = a[local-1][modValue]
				//3,0
				temp4 = a[modValue][i]
			} else {
				temp3 = a[local-1][modValue-1]
				//3,0
				temp4 = a[modValue-1][i]
			}

			//move 3,0 to 0,0
			a[i][j] = temp4
			//move 0,0 to 0,3
			a[j][local-1] = temp1
			if i != 0{
				//move 0,3 to 3,3
				a[local-1][modValue] = temp2
				//move 3,3 to 3,0
				a[modValue][i] = temp3
			} else {
				//move 0,3 to 3,3
				a[local-1][modValue-1] = temp2
				//move 3,3 to 3,0
				a[modValue-1][i] = temp3
			}
			for k := 0; k < arraySize; k++{
				for l := 0; l < arraySize; l++{
					print(" ", a[k][l])
				}
				print("\n")
			}
			fmt.Println("------")
		}
	}

	return a
}

func realRotate(a [][]int) [][]int{

	arraySize := len(a)

	for i := 0; i < arraySize/2; i++{
		for j := i; j < arraySize - i - 1; j++{
			temp := a[j][i]
			a[j][i] = a[arraySize - 1 - i][j]
			a[arraySize - 1 - i][j] = a[arraySize - 1 - j][arraySize - 1 - i]
			a[arraySize - 1 - j][arraySize - 1 - i] = a[i][arraySize - 1 - j]
			a[i][arraySize - 1 - j] = temp
		}
	}
	return a
}

func main(){

	arraySize := 5
	temp := make([][]int, arraySize)
	for i := 0; i < arraySize; i++{
		temp[i] = make([]int, arraySize)
	}

	value := 1

	for i := 0; i < arraySize; i++{
		for j := 0; j < arraySize; j++{
			temp[i][j] = value
			value += 1
		}
	}

	realRotate(temp)

}
