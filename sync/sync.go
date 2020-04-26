package sync

import (
	"sync"
	"fmt"
)

func Once(){
	var once sync.Once
	
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(func(){
				fmt.Println("hello world")
			})
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func WaitGroup(){
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			fmt.Println(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}



func Mutex(){
	var lock sync.Mutex 

	for i := 0 ; i < 10 ; i ++ {

		
		go func(i int) {
			lock.Lock()
			fmt.Printf("%v hello world\n",i)
			lock.Unlock()
		}(i)

		
	} 

}