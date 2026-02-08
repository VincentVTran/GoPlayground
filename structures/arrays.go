package structures

import "fmt"

func DemoArrays() {

	// Declare an array of integers with a fixed size of 5
	var arr [5]int

	// Initialize the array with values
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	// Print the array
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	// Create array list
	arrList := []int{1}
	arrList = append(arrList, 6) // Add an element to the array list
	arrList = append(arrList, 8)
	fmt.Println(arrList)
}
