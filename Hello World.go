package main	//Package is like a project or workspace. It is a collection of common source code files. 
				// Go consists of two different types of packages. The executable and the resuable.
				// Executable generates a file that we can run while reusable uses code as 'helpers', 
				// Good place to put reusable logic. COde dependencies or libraries
				// Main is specifically used to make a executable type package. This MUST have a func called main!

import "fmt"	// Gives the above package all the code and functionality that is contained inside of the other package fmt
				//It is a standard library package

// For rxamples are as below
// import "math"	//gives access to the math library package
// import "debug"
// import "crypto"
// import "encoding"
// import "io"

// We can also import packages that have been created by other engineers i.e reusable packages


func main() {		// Package main should have func main.

	fmt.Println("Hello World!")

}
