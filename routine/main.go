package main

import "fmt"

var pi = 3.1415926


//package scope
func scope() float64{
	var pi = 3.1415
	return pi
}

func loop(){
	for i := 0; i < 10 ;i ++ {
		fmt.Printf("for condition %d\n", i)
	}

	var arr = []byte{1,2,3}

	for index,v := range arr {
		fmt.Printf("for range %d %d\n", index,v)
	} 
}

func main(){
	fmt.Println(pi)

	var pi = scope()
	fmt.Println(pi)

	loop()
	// fmt.Printf(range 10)
}