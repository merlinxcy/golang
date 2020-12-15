package main

import "fmt"
import "time"

func main(){
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	go func() {
		for{
			select {
			case <-done:
				fmt.Println("done")
				return
			case t:= <-ticker.C:
				fmt.Println("ticker at : ",t)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	ticker.Stop()
	done<- true
}
