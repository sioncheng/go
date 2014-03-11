package main

import "fmt"

type ArgumentError struct{
	Message string;
}

func(e *ArgumentError) Error()string {
	return "ArgumentError: " + e.Message;
}

func div(a int,b int)(result int,err error){
	if(b == 0){
		return 0,&ArgumentError{"b can not be zero"};
	}

	return a/b,nil;
}

func printResult(r int,e error) {
	if e == nil {
	    fmt.Printf("%d\n", r);
    }else{
    	fmt.Println(e);
    }
}

func main() {
	r,e := div(4,2);

	printResult(r,e);

	printResult(div(4,0));
}