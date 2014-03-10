package main

import "fmt"

type Person struct {
	FirstName string;
	LastName string;
}

func(e *Person) Hello() string{
	return "My name is " + e.FirstName + " " + e.LastName;
}

func main() {
	var p Person;

	p.FirstName = "Sion";
	p.LastName = "Cheng";

	fmt.Println(p.Hello());
}