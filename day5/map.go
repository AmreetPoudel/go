package main

import (
	"fmt"
)

func main() {
	fmt.Println("map setup done")
	// map is nothing but key value store like dicitonary
	// keys= unique identifier (name age or similar)
	// values= corresponding values for keys
	// simply crating map is easy
	map1 := make(map[string]string)

	map1["name"] = "amrit poudel"
	map1["age"] = "26"
	map1["height"] = "170"
	// user_map := map[string]string{
	// 	"name":   "amrit poudel",
	// 	"age":    "26",  //in years
	// 	"height": "180", //in cm
	// }

	fmt.Println("printing map map1")
	for key, value := range map1 {
		fmt.Printf("%s : %s\n", key, value)
	}
	// fmt.Println("printing user_map")
	// age_in_month, _ := strconv.Atoi(user_map["age"])
	// age_in_month *= 12
	// fmt.Printf("age in month:%d\n", age_in_month)

	// for key, value := range user_map {
	// 	fmt.Printf(" %s = %s\n", key, value)

	// }

	// fmt.Println("deleting the data from map ")
	// delete(map1, "height")
	// fmt.Println(map1)
	// fmt.Println("deleting everything from map")
	// clear(map1)
	// fmt.Println(map1)
	name, does_data_exists := map1["name"]
	println(does_data_exists, name)

}
