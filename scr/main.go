/*
	Summary

	Arrays
		Collection of items with same type
		Fixed size
		Declaration
			a := [3]int{1, 2, 3}
			a := [...]int{1, 2, 3}
			var a [3]int
		Access via zero-based index
			a := [3]int{1, 3, 5} // a[1]==3
		len function returns size of array
		Copies refe to different underlying data

	Slices
		Backed by array
		Creation styles
			Slice existing array or slice
			Literal style
			Via make function
				a := make([]int, 10) // create slice with capacity and lenght == 10
				a := make([]int, 10, 100) // slice with lenght == 10 and capacity == 100
		len function returns lenght of slice
		cap function return lenght of underlying array
		append function to add elements to slice
			May cause expensive copy operation if underlying array is too small
		Copies refer to same underlying array


*/

package main

import "fmt"

func main() {
	// to create an array
	grades := [3]int{97, 85, 93} // [97, 85, 93], len = 3
	// if we want to create array with amount of elements just
	// to store initialised values then just do
	// [...]int{91, 85, 93} len = 3!
	fmt.Printf("Grades: %v\n", grades) // Grades: [97 85 93]

	// also we can declare array like this
	var students [3]string                 // [ ] empty, len = 3
	fmt.Printf("Students: %v\n", students) // Students: [ ]
	// if we want to assign to first element of array a name we use
	students[0] = "Ivan Kosovan"                   // Ivan Kosovan
	fmt.Printf("First student: %v\n", students[0]) // First student: Ivan Kosovan
	// to get the size of array simply do len(arrayName)
	fmt.Printf("Ammount of students: %v\n", len(students))
	// & - is a pointer operator
	// to create a slice
	slice := []int{1, 2, 3}
	fmt.Printf("%v\n", slice)
	// we can use len for slices too, but also we can use
	// a cap which will show us how much elements are not null
	// to append an element to slice you need to use append(slice, value) function
	// also you can append not only one element to slice, you can do it like that
	// append(slice, value, value, value ,...)
	slice = append(slice, 228)
	fmt.Printf("%v\n", slice)
	// to remove the first element from the slice
	// we use b:=slice[1:] [2, 3, 228]
	// to pop the last element we need to do
	//	b:=slice[:len(slice)-1] [1, 2, 3]
}
