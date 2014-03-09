package main

import "fmt"

func main(){
	
	var int_arr [10]int = [10]int{1,2,3,4,5,6,7,8,9,0};

	fmt.Println("the first element of int_arr");
	fmt.Println(int_arr[0]);

	fmt.Println("modify the value of first element of int_arr");
	int_arr[0] = 100;
	fmt.Println(int_arr[0]);

	fmt.Println("go through all elements of int_arr");
	for i,v := range int_arr{
		fmt.Printf("%d, %d\n", i,v);
	}

	fmt.Println("make a slice of int_arr by cutting last 5 elements");
	int_arr_slice := int_arr[5:]
	for i,v := range int_arr_slice{
		fmt.Printf("%d, %d\n", i, v);
	}

	fmt.Println("add new element to slice");
	int_arr_slice = append(int_arr_slice,100)
	fmt.Println(int_arr_slice[5]);
	fmt.Printf("len %d, cap %d\n", len(int_arr_slice), cap(int_arr_slice));

	fmt.Println("copy elements from int_arr to int_arr_slice");
	copy(int_arr_slice,int_arr[0:]);
	for i,v := range int_arr_slice{
		fmt.Printf("%d, %d\n",i,v);
	}
}