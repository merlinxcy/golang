//package main
//
//import (
//	"fmt"
//	"time"
//)
//func get11(p1 string){
//	for{
//		time.Sleep(time.Second)
//		fmt.Println(p1)
//	}
//
//}
////go Get11(p1,p2)
//func main(){
//	var input string
//	go get11("12")
//	fmt.Println("sad")
//	fmt.Scanf("%s",&input)
//	fmt.Println(input)
//	fmt.Println("sad")
//	time.Sleep(time.Minute)
//
//}

//package main
//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//var counter int = 0
//func Count(lock *sync.Mutex){
//	lock.Lock()
//	counter++
//	fmt.Printf("in count: %d\n", counter)
//	lock.Unlock()
//}
//func main(){
//	lock := &sync.Mutex{}
//	for i:=0;i<10; i++{
//		go Count(lock)
//	}
//	a := make(map[int]string)
//	a[1] = "das"
//	for i:= range a{
//		fmt.Println(i)
//	}
//	for{
//		lock.Lock()
//		c := counter
//		fmt.Println(c)
//		lock.Unlock()
//		runtime.Gosched()
//		if c>=10{
//			break
//		}
//	}
//}

//package main
//import (
//	"fmt"
//)
//
//type DataWrite interface {
//	WriteData(data interface{}) error
//}
//
//type file struct {
//
//}
//
//func (d *file) WriteData(data interface{}) error{
//	fmt.Println("WriteData:", data)
//	return nil
//}
//func main(){
//	f := new(file)
//	var write DataWrite
//	write = f
//	write.WriteData("data")
//
//}

//package main
//
//import "awesomeProject/base"
//import _"awesomeProject/cls"
//
//func main(){
//	cls := base.Create("cls1")
//	cls.DO()
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main(){
//	ch := make(chan int)
//	go func(){
//		fmt.Println("start goroutine")
//		ch <- 0
//		fmt.Println("exit goroutine")
//	}()
//	//time.Sleep(1000)
//	fmt.Println("wait goroutine")
//	<-ch
//	fmt.Println("all done")
//}

//package main
//
//import "fmt"
//
//func printer(c chan int){
//	for{
//		data := <- c
//		if data == 0{
//			break
//		}
//		fmt.Println(data)
//
//	}
//	c <- 0
//}
//
//func main(){
//	c := make(chan int)
//	go printer(c)
//	for i :=1;i<=10; i++{
//		c <-i
//	}
//	c<-0
//	<-c
//}

//package main
//
//import "fmt"
//
//func main(){
//	demo := map[string]bool{
//		"a": false,
//	}
//	fmt.Println(demo["a"])
//	_, ok := demo["a"] //判断a是否存在
//	fmt.Println(ok) // true
//	_, ok = demo["b"] //判断b是否存在
//	fmt.Println(ok) // false
//}

//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//func main(){
//	resp, _ := http.Get("https://xeldax.top")
//	fmt.Println(resp.StatusCode)
//	len := 4096
//	fmt.Println(len)
//	buf := make([]byte,len)
//	resp.Body.Read(buf)
//	fmt.Println(string(buf))
//
//	}
//
//package main
//
//import (
//	"fmt"
//	"net"
//	"time"
//)
//func checkonporint(ip string,i int){
//	target := fmt.Sprintf("%s:%d",ip,i)
//	_,err := net.DialTimeout("tcp4", target,time.Millisecond * 1000)
//	//defer handle.Close()
//	if err != nil{
//		//fmt.Println(err)
//		return
//	}
//	fmt.Printf("Port is open: %d\n", i)
//	return
//}
//func main(){
//	//wait := make(chan int)
//	var ip string
//	ip = "122.226.73.152"
//	//flag.StringVar(&ip, "i", "", "ip address")
//	//flag.Parse()
//	fmt.Printf("Scaning: %s\n", ip)
//	for i:=1;i<=65535;i++{
//		//fmt.Printf("12")
//		go checkonporint(ip,i)
//	}
//	time.Sleep(time.Minute)
//}
