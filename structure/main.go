package main

func main(){
	
}


func scope(){
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c\n", x) // "HELLO" (one letter per iteration)
		}
	}
}