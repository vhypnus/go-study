package main 


import "fmt"

func (n int)  add(j int) int {

	return n+j 
}

func main(){
	var i = new(3)
	fmt.Println(i.add(6))
}