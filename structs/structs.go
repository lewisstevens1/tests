package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person2 struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)

	// jim := person2{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		email:   "lewis@lewis.com",
	// 		zipCode: 96305,
	// 	},
	// }
	// fmt.Printf("%+v", jim)

	dave := person{
		firstName: "Dave",
		lastName:  "Tool",
		contactInfo: contactInfo{
			email:   "dave@dave.com",
			zipCode: 96305,
		},
	}

	// You need to assign a pointer as calling a new function will cause a new address in memory to be allocated. This way we can point back to the old value and update that.
	// davePointer := &dave
	// davePointer.updateName("OOPS")
	dave.updateName("aaa")
	dave.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// &variable - give me mem address of this var.
// *pointer
// You can skip passing the memory address because we are using * before type in the reciever

// * before type (*person) - tells it that it should be a person.
// * before pointer (*pointerToPerson) - take this mem address and give me the value.

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
