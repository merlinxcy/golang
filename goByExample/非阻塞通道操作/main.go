package main

import (
	"fmt"
	"time"
)

func main(){
	messages := make(chan string)
	signal := make(chan string)
	//test := make(chan bool)
	//go func() {
	//	time.Sleep(time.Second * 100)
	//	test <- true
	//}()
	//<-test
	select {
	case <-messages:
		fmt.Println("recevied message")
	default:
		fmt.Println("no message")
	}

	go func() {
		signal <- "signal"
		fmt.Println("asdad")
	}()
	time.Sleep(time.Second)
	select {
	case <- signal:
		fmt.Println("recevied signal")
	default:
		fmt.Println("no signal")
	}

	select {
	case <-signal:
		fmt.Println("signal")
	case <-messages:
			fmt.Println("acccc")
	default:
		fmt.Println("no acitive")
	}
	//非阻塞发送
	msg := "hihihiihhi"
	select {
	case messages <- msg:
		fmt.Println("send message")
	default:
		fmt.Println("no message sent")
	}
}
