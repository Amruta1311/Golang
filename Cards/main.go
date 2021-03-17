package main

func main() {

	//GoLang is not an OOP Language so there is no concept of classes here

	// card := newCard()

	// := is only used when we are initialising the variable for the first time. Wo't work if we use it for reassignment. Use equal (=) there

	// var card string = "Ace of Spades"	// Variable declaration of string type and its assignment (Other data types are byte, bool, int, float64)

	// Alternative way of writing the above line of code is as follows

	//============REAL CODE============//

	// cards := []string{"Ace Of Diamonds", newCard()} // SLICE data structure of Go, SLICE can contain a group of same data type within it

	// Introducing deck variable as an extensible of slice data structure helps us use the same in main.go

	// cards := deck{"Ace Of Diamonds", newCard()}

	// Adding an additional card to the above cards

	// cards = append(cards, "Six Of Spades")

	//SLICE OF BYTES: Predominantly type converts a given string into its ascii form of arrays/slices, e.g. []byte("Hello")

	//----------------------------
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// fmt.Println("Hand of Cards:")

	// hand.print()

	// fmt.Println("Remaining Cards:")

	// remainingCards.print()

	//----------------------------

	//cards := newDeckFromFile("My_Cards")

	cards := newDeck()

	cards.saveToFile("My_Cards") //Saved to file using the filename 'My_Cards' in our Hard Drive

	cards.shuffle()

	cards.print()

	// fmt.Println(cards.toString())

	// cards.print() // Any variable of type deck gets access to the "print" method

	// To obtain the deck of cards we are going to extend the base class and add some extra functionalities
	// to it. We tell go to create an array of strings and add a bunch of functions specifically made to work with it.

}

// func newCard() string { //We tell go compiler that what type of data are we returning to the function

// 	return "Five Of Diamonds"

// }
