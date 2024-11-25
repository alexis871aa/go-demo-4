package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a) // меняет порядок элементов на обратный
	fmt.Println(a)
}

func reverse(arr *[4]int) {
	// другой вариант (якобы правильный)
	for index, value := range *arr {
		(*arr)[len(arr)-1-index] = value
	}

	// мой вариант
	//for i := 0; i < len(*arr)/2; i++ {
	//	tmp := (*arr)[i]
	//	(*arr)[i] = (*arr)[len(*arr)-i-1]
	//	(*arr)[len(*arr)-i-1] = tmp
	//}
}
