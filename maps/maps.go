package maps

/*
	In Go, Maps is an unordered collection of key and its associated value. They are very good for looking up values fast.
	The key type must have the operations == and != defined, like string, int, float

	Hence arrays, slices and structs cannot be used as key type, but pointer and interface types can

	structs can be used as a key when we provide key() or Hash method, so that a unique numeric or string key can be calculated from the struct's fields.

	Declared as: var map1 map[keyType] valueType
	E.g., var map1 map[int] string
*/
import (
	"fmt"
)

type Vertex struct {
	Latitude, Longitude float64
}

func Maps() {
	// Go Map initialization
	x := map[string]string{"firstName": "Sachin", "lastName": "Yadav"}
	fmt.Println(x)
	fmt.Println("Only FirstName: ", x["firstName"])

	// Go map insert, retrieve, update and delete
	myMap := make(map[string]int)
	fmt.Println("Empty Map: ", myMap)

	// insert
	myMap["key1"] = 10
	myMap["key2"] = 20
	myMap["key3"] = 30
	fmt.Println(myMap)
	// update
	myMap["key2"] = 50
	fmt.Println(myMap)

	// delete
	delete(myMap, "key1")
	fmt.Println(myMap)

	// retrieve
	var elem = myMap["key3"]
	fmt.Println(elem)

	// to test key is present or not
	var ok bool
	elem, ok = myMap["key2"]
	fmt.Println("The value:", elem, "Present?", ok)

	// Map of struct
	var m = map[string]Vertex{
		"JavaPoint": Vertex{40.68433, -74.39967},
		"SSS-IT":    Vertex{37.42202, -122.08408},
	}
	fmt.Println(m)
}
