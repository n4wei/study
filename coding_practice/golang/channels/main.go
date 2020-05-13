package main

import (
	"fmt"
	"time"
)

func main() {
	doneCh := make(chan struct{})
	dataCh := make(chan string, 100)

	go process1(dataCh, doneCh)
	go process2(dataCh, doneCh)
	go signalStop(doneCh)

	loop:
	for {
		select {
		case data, ok := <-dataCh:
			if !ok {
				break loop
			}
			fmt.Println("received: ", data)
		}
	}

	<-doneCh
	fmt.Println("exit main")
}

func signalStop(doneCh chan struct{}) {
	time.Sleep(3*time.Second)
	close(doneCh)
}

func process1(dataCh chan string, doneCh chan struct{}) {
	defer func() {
		fmt.Println("exit process1")
		close(dataCh)
	}()

	ticker := time.NewTicker(100*time.Millisecond)
	for {
		select {
		case <-ticker.C:
			dataCh<-"process1"
		case <-doneCh:
			ticker.Stop()
			return
		}
	}
}

func process2(dataCh chan string, doneCh chan struct{}) {
	defer func() {
		fmt.Println("exit process2")
	}()

	ticker := time.NewTicker(500*time.Millisecond)
	for {
		select {
		case <-ticker.C:
			dataCh<-"process2"
		case <-doneCh:
			ticker.Stop()
			return
		}
	}
}