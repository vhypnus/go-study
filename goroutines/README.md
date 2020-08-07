
```go
f() //call f();wait it for return

go f() //create a new goroutine that calls f();don't wait

```

### channel
```
var ch = make(chan int) //ch has type 'chan int',and unbuffer channel
var ch = make(chan int,0) //also unbuffer channel
var ch = make(chan int,3) //buffered channel with capacity 3

```

- A channel has principal operations : send and receive,collectively known as communications.
- A send statement transimits a value from one goroutine ,through the channel,to another goroutine executing a corresponding receive expression.

```
ch <- x // a send statement
x = x <- ch // a receive expression in an assignment statement

<- ch // a receive ecpression,result is discarded
``` 

### unbuffer channel 
- synchronize

### Pipelines