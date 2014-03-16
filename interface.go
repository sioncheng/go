package main

import "fmt"

type IHello interface{
	Hello()string;
}


type Person struct {
	FirstName string;
	LastName string;
}

func(e *Person) Hello() string{
	return "My name is " + e.FirstName + " " + e.LastName;
}

func main() {
	var ihello IHello;

	ihello = &Person{FirstName:"Sion",LastName:"Cheng"};

	fmt.Println(ihello.Hello());
}