package basics

import "fmt"

func main() {
	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 19
	fmt.Println(myMap)
	fmt.Println(myMap["code"])
	delete(myMap, "key1")
	fmt.Println(myMap)
	clear(myMap)
	fmt.Println(myMap)

	var myMap3 map[string]string
	myMap3 = make(map[string]string)
	myMap3["key"] = "Value"
	fmt.Println(myMap3)
	myMap4 := make(map[string]map[string]string)

	myMap4["map1"] = myMap3
	fmt.Println(myMap4)

}
