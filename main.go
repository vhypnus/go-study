package main


import (
	"fmt"
)

func main(){

	var i = new(int)

	fmt.Println(&i)
	*i = 3

	fmt.Println(&i)

	var v interface{}

	v = &Sample{10}
	fmt.Println(v)

	var x interface{} = 7
	y := x.(int)
	fmt.Println(y)

	var s1 = &Sample{7}
	var s2 = &Sample{7}
	fmt.Println(*s1)
	fmt.Println(*s2)
	fmt.Println(*s1 == *s2)
}


type Sample struct {
	i int 
}