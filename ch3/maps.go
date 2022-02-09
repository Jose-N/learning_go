package main

import "fmt"

func main() {

	// Instantiating maps and display how non-existant key returns 0 value (default value for int)
	fmt.Println("READING AND WRITING A MAP")
	totalwins := map[string]int{}
	totalwins["Orcas"] = 1
	totalwins["Lions"] = 2
	fmt.Println(totalwins["Orcas"])
	fmt.Println(totalwins["Kittens"])
	totalwins["Kittens"]++
	fmt.Println(totalwins["Kittens"])
	totalwins["Lions"] = 3
	fmt.Println(totalwins["Lions"])

	// Use two variable assignment retiving key from map to get both value and bool if exists in map
	fmt.Println("COMMA OK IDIOM")
	m := map[string]int{
		"hello":   5,
		"goodbye": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	fmt.Println("DELETING FROM MAPS")
	delete(m, "hello")
	fmt.Println(m)
}
