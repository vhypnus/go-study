// Interface types express generalizations and abstraction about the behaviors of other types.
// By generalizing interfaces let us write functions that are more flexible and adaptable because 
// they are not tied to the details of one particular implemention.

/*

package io

type Reader interface {
	Read(n []byte) (n int,err error)
}

type Closer interface {
	Close()	error
}

type Writer interface {
	Write(n []byte) error
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}



*/

package main

import (
	// "io"
	"time"
)


// In Go,varirable are always initialized to a well-defined value,and interfaces are no exception.	
func Interfaces(){

	// var w io.Writer
	// w.Write([]byte("hello")) //panic:nil pointer dereference

	//an interface value can hold arbitrarily large dynamic values. 
	var x interface{} = time.Now()

	// if the type is comparable ,two interface has the same dynamic type

	//type assertions

	// var w io.Writer
	// w = os.Stdout
	// f := w.(*os.File) //Success: f == os.Stdout
	// c := w.(*bytes.Buffer) //panic : interface hold *os.File,not *bytes.Buffer


	// First, if the ass erted typ e T is a con crete typ e, then the typ e assertion che cks whether x’s dynamic typ e is ident ical to T.
	// If this che ck succe e ds, the result of the
	// type ass ertion is x’s dynamic value, whose typ e is of cours e T.

	//Second, if ins tead the ass erted typ e T is an int erface typ e, then the typ e assertion checks
	//whether x’s dynamic typ e sati sfies T. If this che ck succe e ds, the dynamic value is not ext racte d;


	// type switch
}


func main(){
	
}
