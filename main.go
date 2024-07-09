package main

import "fmt"

func main() {

// Declaring Variables
// There are Four Methods for Declaring Variables in GoLang

var nameOne string = "Fabian";
var nameTwo = "Amino";
var nameThree string;

fmt.Println(nameOne, nameTwo, nameThree)

nameOne = "Harry"
nameThree = "Faya"

fmt.Println(nameOne, nameTwo, nameThree)

// This is the shorthand Method
nameFour := "Keziah"

fmt.Println(nameOne, nameTwo, nameThree, nameFour)

// Declaring Integers


var ageOne int = 35;
var ageTwo = 30;
ageThree := 40;


fmt.Println(ageOne, ageTwo, ageThree)


//bits and Memory

var numOne int8 = 25;
var numTwo int8 = -120;
var numThree uint8 = 205;

fmt.Println(numOne, numTwo, numThree)

// formatting strings

fmt.Print("hello, \n" )
fmt.Print("World, \n" )
fmt.Print("new line")


// Println
fmt.Println("Hellow People")
fmt.Println("My name is", nameOne, "and I am", ageOne, "years old")


//Printing Formated Strings
fmt.Printf("my name is %v and I am %v years old \n", nameTwo, ageTwo)
fmt.Printf("my name is %q and I am %q years old \n", nameTwo, ageTwo)
fmt.Printf("age is of type %T \n", ageTwo)
fmt.Printf("you scored %0.1f points \n", 5666.888)


// Sprintf (save formatted strings)
var str = fmt.Sprintf("my name is %v and I am %v years old \n", nameTwo, ageTwo)
fmt.Println("the saved string is: ", str)




}
