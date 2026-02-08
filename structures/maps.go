package structures

import "fmt"

func DemoMaps() {

	// Declare a map with string keys and integer values
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 10
	myMap["banana"] = 20
	myMap["orange"] = 30

	// Print the map
	for key, value := range myMap {
		println(key, value)
	}

	// Check if a key exists in the map
	if value, exist := myMap["banana"]; exist == true {
		fmt.Println("Value for 'banana':", value)
	} else {
		fmt.Println("'banana' does not exist in the map")
	}

	// Delete a key-value pair from the map
	delete(myMap, "apple")

	// Print the map after deletion
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
