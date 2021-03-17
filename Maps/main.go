package main

import (
	"fmt"
)

func main() {
	//First way of declaring a map

	color := map[string]string{ //Here we declare a map where the key is of type string and the values are also of type string

		"red": "#ff0000",

		"green": "#00ff00",

		"blue": "#0000ff",

		"black": "#1324de",

		"indigo": "#8923a7",

		"yellow": "#54d2e3",
	}

	fmt.Println("First: \n ", color)

	//Second way of declaring a map: If we do not assign a value to the map then go will initialise it to zero

	var colors map[string]string

	fmt.Println("Second: \n ", colors)

	//Third way of declaring a map

	col := make(map[string]string) //make() is a built in function. This creates an empty map too

	col["white"] = "#01234f"

	col["orange"] = "#ab23c5"

	fmt.Println("THird: \n ", col)

	//Deleteing a key out of the map

	delete(color, "green") // delete is an inbuilt function. Here we pass the map name and the key name to delete the particular entry from the list

	fmt.Println("After deleteing green in First: \n", color)

	//Iterating Over a Map

	printMap(color)
}

func printMap(c map[string]string) {
	fmt.Println("Inside the printMap function:")
	for key, value := range c {
		fmt.Println("Hex code for ", key, " is ", value)
	}
}
