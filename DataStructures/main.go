package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastname  string
	contact   contactInfo // Alternatively you may write, contactInfo, which defines a variable contactInfo of type contactInfo
} //Defining a new custom type person of struct that will only exist inside of the program

func main() {
	// amu := person{"Amruta", "Mulay"}

	// Other way of creating a new person is as follows

	amu := person{
		firstName: "Amruta",
		lastname:  "Mulay",
		contact: contactInfo{ // If followed the alternative given above then here, contactInfo: contactInfo, is written
			email:   "amruta13mulay11@gmail.com",
			zipCode: 940000,
		},
	}

	//Last way of declaring a struct is as follows

	var Amu person //No properties assigned yet. GO assigns a zero value to the ones that are not yet assigned. In our case it is empty string

	fmt.Printf("%+v", amu) //Used as an interpolation syntax where we put an identifier that then gets filled in with an actual value

	fmt.Println("\n", Amu)

	//In the above statement, %+v prints out all the different names and their values from Amu

	//Updating the values or fields on the struct

	Amu.firstName = "AMRUTA"
	Amu.lastname = "MULAY"

	fmt.Printf("\n Amu: %+v", Amu)

	// fmt.Println("\n ======================:Printing using a reciever:============================")

	// amu.print()

	// amu.updateName("AMY")

	fmt.Println("\n ======================:Printing using a reciever:============================")

	amu.print() //Update did not take effect in the application. This is where pointers come to picture

	//Go is referred as the pass by value language. Whenever we pass by value, golang takes that value or struct and copy all that data inside that struct and place in it some new container of the memory.
	// The reciever p is pointing to thsi new container and makes changes with respect to this new container.
	// Thus when we run the print function, it printed the original container

	amuPointer := &amu //Gives the memory address of the value amu is pointing to

	amuPointer.updateName("AMY")

	// The above two lines can also be written as: amu.updateName("AMY")
	// It will automatically convert the person type to the pointer to person type before passing it as the reciever to the function

	fmt.Println("\n ======================:Using the Pointers:============================")

	amu.print()
}

// *person signifies a description to a type. We are looking for a pointer to a person. updateName function can only be called with the receiver of a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) { //*pointerToPerson gives the value that this memory address is pointing to
	// The below is a operator - it means we want to manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
