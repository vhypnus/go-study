package main

import (
	"fmt"
	// "sort"
)

type Currency int

const(
	USD Currency = iota
	EUR
	GBP
	RMB
)


func Array(){
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	//print index and value 
	for i,v := range a {
		fmt.Printf("print index and value %d %d \n",i,v)
	}

	//print value only
	for _,v := range a {
		fmt.Printf("print value only %d\n",v)
	}


	//
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	// if an ellipsis ‘‘...’’ appears in place of the lengt h, the array lengt h is determined by the number of initializers. 
	q = [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"


	q = [3]int{1, 2, 3}
	// q=[4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int


	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: "￥"}

	// symbol := [...]string{3: "$", EUR: "9", GBP: "!", 0: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ""


	// zero value
	var months [12]int
	fmt.Println(months)
}


/* 
* (1)slice writted by []T.
* (2)a slice has three compnent a pointer,a length and a capacity.
* (3)A slice is a lightweight data structure tha gives access to a subsequence of the elements of an array,
	which is known as the slice's underlying type.
  (4) slice are not comparable ,so we cannot use == to test where two slice contain the sam elements.But bytes.Equal function 
    can comparable two slice  of bytes.
**/
func Slice(){
	var months = []string{"January","February","March","April","May","June","July","August","September","October","November","December"}
	var q2 = months[3:6]
	fmt.Println(q2)

	//if you need test whether a slice is empty,use len(s)==0,not s == nil.
	var s []string
	fmt.Println(s)
	if s == nil {
		fmt.Println("it will not happen.")
	}
}

// map is a hash table  is written by map[k]
// the built -in function make can be used to create a map.
func Maps(){

	// var ages = make(map[string]int)
	// ages["huangyuewen"] = 35
	// ages["zhangfenlan"]= 36

	var ages =map[string]int{
		"huangyuewen":35
		"zhangfenlan":36
	}

	// delete(ages,"zhangfenlan")

	//but map element is not a variable ,and we can't take it's address.
	// _ = &ages["huangyuewen"] //compile error: can't take address of map element


	//if you get a key not exists in map

	age,ok := ages["bob"]
	if !ok{
		fmt.Println("bob not born")
	}


}


// a struct is a aggregate data type that groups together zero or more named values of arbitrary types of a single entity
// each value is called a field.

// type Point struct{x,y int} //x,y are not exported

type Point struct{X,Y int}

//embedding 

type Circle struct {
	// center Point
	Point
	radius int
}

type Wheel struct{
	// Circle Circle
	Circle
	Spokes int 
}

func Struct(){
	var p := Point(1,2)
	fmt.Println(p)

	// not exported demo

	/*
	package p
	type T struct{ a, b int } // a and b are not exported
	package q
	import "p"
	var _ = p.T{a: 1, b: 2} // compile error: can't reference a, b
	var _ = p.T{1, 2} // compile error: can't reference a, b
	*/
}


func Json(){
	
}

func Reverse(){

}

func main(){

	Array()
	// Slice()

}