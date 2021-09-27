package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) []int {
	var newslice []int
	for _, val := range slice {
		newslice = append(newslice, f(val))
	}
	return newslice
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	var newarr [3]int
	for i, val := range array {
		newarr[i] = f(val)
	}
	return newarr
}

func main() {
	intsSlice := []int{1, 2, 3}
	intsSlice = mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [3]int{1, 2, 3}
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)

	newSlice := intsSlice[1:3]
	newSlice = mapSlice(square, newSlice)
	fmt.Println(newSlice)

	// arrays are passed by value; slices are passed by reference
	// (well, by value of a struct with length and a pointer to the start)

}
