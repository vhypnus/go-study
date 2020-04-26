package sync 

import (
	"testing"
	"time"
)


func TestOnce(t *testing.T){
	Once()

	time.Sleep(1*time.Second)
}

func TestWaitGroup(t *testing.T) {
	WaitGroup()
}

func TestMutex(t *testing.T) {
	Mutex()
	time.Sleep(3*time.Second)
}