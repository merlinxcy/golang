package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "wait for time one second"
	}()

	select {
	case data := <- c1:
		fmt.Println(data)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "wait for time one second"
	}()
	select{
	case data := <- c2:
		fmt.Println(data)
	case <-time.After(time.Second * 10):
		fmt.Println("time out")

	}
}
