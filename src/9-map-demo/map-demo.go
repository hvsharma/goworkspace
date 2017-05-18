package main

import (
	"fmt"
	"sort"
)

func main() {
	statesMap := make(map[string]string)
	statesMap["TS"] = "Telangana State"
	statesMap["MP"] = "Madhya Pradesh"
	statesMap["RJ"] = "Rajasthan"
	statesMap["UP"] = "Uttar Pradesh"
	statesMap["AP"] = "Andhra Pradesh"

	fmt.Println(statesMap)

	mp := statesMap["MP"]
	fmt.Println("MP:", mp)

	delete(statesMap, "RJ")
	fmt.Println("Map after deleting RJ:", statesMap)

	for k, v := range statesMap {
		fmt.Printf("%v - %v\n", k, v)
	}

	size := len(statesMap)
	i := 0
	keys := make([]string, size)
	for k := range statesMap {
		keys[i] = k
		i++
	}

	fmt.Println("Keys -:", keys)
	sort.Strings(keys)
	fmt.Println("Sorted Keys -:", keys)
	fmt.Println("Sorted Map -:")

	for x := range keys {
		fmt.Printf("%v - %v\n", keys[x], statesMap[keys[x]])
	}

}
