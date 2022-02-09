package main

import "fmt"

func main() {
	// structs defined in functions can only be used by that function
	type person struct {
		name string
		age  int
		pet  string
	}

	// variables can be define via type or by a struct literal
	// both below cases assign zero vaule to each field in the struct
	var fred person
	bob := person{}

	fmt.Println("EMPTY STRUCTS")
	fmt.Println(fred)
	fmt.Println(bob)

	// below are two styles of declaring nonempty structs literal
	// fields must be assigned in the order of the struct def, fine for small structs
	julia := person{
		"Julia",
		40,
		"cat,",
	}

	// this allows assignement of field in any order, any left out fields assigned to zero value
	// this is clearer and more mantainable
	beth := person{
		age:  30,
		name: "Beth",
	}

	// setting fields in struct
	bob.name = "Bob"
	fred.age = 25

	// accessing field in struct
	fmt.Println("Bob's name is " + bob.name)
	fmt.Println("Freds age is " + fmt.Sprint(fred.age))

	fmt.Println(julia)
	fmt.Println(beth)

	// anonymous structs are also a thing, var instead of type
	// useful json un/marshalling and tests
	var pet struct {
		name string
		kind string
	}

	pet.name = "Fido"
	pet.kind = "Dog"
	fmt.Println(pet)

	dog := struct {
		name  string
		breed string
	}{
		name:  "Bobo",
		breed: "Husky",
	}
	fmt.Println(dog)
}
