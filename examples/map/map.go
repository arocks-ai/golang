package main

import (
	"fmt"
)


func main() {

	fmt.Println("== Empty Map example ...") 
	month := make(map[string]int)	// make is not need if empty map
	month["Jan"] = 1
	month["Feb"] = 2
	month["Mar"] = 3
	month["Apr"] = 4

	fmt.Println("Get length using len(month) = ", len(month)) 	// length
	fmt.Println("Get value using month[\"Feb\"] = ", month["Feb"]) 		// get value	
	fmt.Println("Key not found example using month[\"RandomString\"] = ", month["RandomString"]) 	// Prints default value if key not found		

	
	fmt.Println("\n==Printing map keys and values = \n", month)	//prints map[k1:v1 k2:v2] format

	fmt.Println("Printing map key-values using range ..")
	for key1, val1 := range month{
		fmt.Println("key = ", key1, "value = ", val1)
	}

	delete(month, "Mar")	// remove one key 
	// delete(month)  //this is not supported use clear() method
	fmt.Println("\nupdated map length after key removal using delete() = ", len(month)) 	// updated length

	_, present := month["Apr"]		// Check if key present without getting value
	fmt.Println("Check key without getting value = ", present)

	clear(month)	// Remove all keys
	fmt.Println("updated map length after remove All using clear() = ", len(month)) 	// updated length


	// make is not needed to create non empty map
	fmt.Println("\n== Non empty Map example ...") 
	weekday1 := map[string]int {
			"Sunday" : 1, 
			"Monday" : 2,
			"Tuesday" : 3,
			}	
	fmt.Println(weekday1["Sunday"])


	// map with value could be any type
	fmt.Println("\n== Map with value could be any type ...") 
	mapStudent := map[string]interface{}{
		"name": "student1",
		"age": 40,
		"isAdmin" : true,
		"marks" : []int {44,55,66,77},
	}

	fmt.Println("mapStudent = ", mapStudent)
	for key1, val1 := range mapStudent {
		fmt.Println("key1 = ", key1, "val1 = ", val1)
	}




}