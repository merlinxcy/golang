//package main
//
//import "fmt"
//
//func getdata() (int, int){
//	return 100, 200
//}
//func tets(){
//	a, _ := getdata()
//	fmt.Println(a)
//}
//func main(){
//	fmt.Println("Hello World!")
//	var a_var int = 100
//	var b_var int = 300
//	a_var = a_var ^ b_var
//	b_var = b_var ^ a_var
//	a_var = a_var ^ b_var
//	fmt.Println(a_var,b_var)
//	fmt.Println(sum(100,300))
//}
//
//func sum(a int, b int) int {
//	var num int;
//	num = a + b;
//	return num;
//}

//package main
//import (
//	//"fmt"
//	"image"
//	"image/color"
//	"os"
//	"math"
//	"log"
//	"image/png"
//)
//func main(){
//	const size = 300
//	pic := image.NewGray(image.Rect(0,0,size,size))
//	for x:= 0; x< size; x++ {
//		for y:=0; y<size; y++{
//			pic.SetGray(x,y,color.Gray{255})
//		}
//	}
//	for x:=0; x<size; x++{
//		s:= float64(x) * 2 *math.Pi/size
//		y := size/2 - math.Sin(s) * size/2
//		pic.SetGray(x, int(y), color.Gray{0})
//	}
//	file,  err := os.Create("sin.png")
//	if err != nil{
//		log.Fatal(err)
//	}
//	png.Encode(file, pic)
//	file.Close()
//}

//package  main
//import (
//	"fmt"
//	"encoding/base64"
//)
//
//func main(){
//	message := "dasdasa"
//	encodeMessage := base64.StdEncoding.EncodeToString([] byte(message))
//	fmt.Println(encodeMessage)
//	data,_ := base64.StdEncoding.DecodeString(encodeMessage)
//	fmt.Println(string(data))
//}

//package main
//import (
//	"strings"
//)
//func main(){
//	a := "asdsd"
//	b := "asdsd"
//	//strings.Builder{}
//}

//package main
//import (
//	"fmt"
//	"strings"
//)
//func main(){
//	a := "hello dssad"
//	location := strings.Index(a, "ds")
//	fmt.Println(a[1:])
//	fmt.Println(a[0:1])
//	fmt.Println(location)
//}

//package main
//import (
//	"fmt"
//)
//
//func main() {
//a := "adsaadasd"
//for i, j := range a{
//fmt.Println(i, string(j))
//}
//}

//package main
//import (
//	"fmt"
//	"math"
//)
//func main(){
//	fmt.Print("int8 range: ",math.MinInt8, math.MaxInt8)
//	fmt.Print("int16 range: " , math.MinInt16,math.MaxInt16)
//	fmt.Print("int64 range: ", math.MinInt64, math.MaxInt64)
//}


//package main
//import (
//	"strconv"
//	"fmt"
//)
//func main(){
//	num := 100
//	str := strconv.Itoa(num)
//	fmt.Println(str)
//	str1 := "110"
//	str2 := "s100"
//	num1, err := strconv.Atoi(str1)
//	if err != nil{
//		fmt.Println(err)
//	}	else {
//		fmt.Println(num1)
//	}
//	num2, err := strconv.Atoi(str2)
//	if err != nil{
//		fmt.Println(err)
//	}	else{
//		fmt.Println(num2)
//	}
//}

//package main
//import (
//	"fmt"
//)
//func main(){
//	var q [3]int = [3]int{1,2,3}
//	fmt.Println(q)
//	for i,j := range q{
//		fmt.Println(i,j)
//	}
//}

//package main
//import (
//	"fmt"
//)
//func main() {
//	var highRiseBuilding [30]int
//	for i := 0; i < 30; i++ {
//		highRiseBuilding[i] = i + 1
//	}
//
//}
//
//package main
//import (
//	"fmt"
//)
//
//func main(){
//	var a []int
//	a = append(a,1,2,3,4)
//	fmt.Println(a)
//}

//package main
//import "fmt"
//func main(){
//	const elementCount = 1000
//	srcData := make([]int, elementCount)
//	for i:=0; i < elementCount; i++{
//		srcData[i] = i
//	}
//	refData := srcData
//	copyData := make([]int, elementCount)
//	copy(copyData, srcData)
//	fmt.Println(refData)
//
//}

//package main
//import "fmt"
//func main(){
//	aaa := []int{1,2,3}
//	bbb := []int{4,5,6}
//	aaa = append(aaa,bbb...)
//	fmt.Println(aaa)
//}

//package main
//import "fmt"
////这个很有用
//func main(){
//	seq := []int{1,2,3,4,5}
//	fmt.Println("oring",seq)
//	index := 2
//	fmt.Println(seq[:index],seq[index+1:])
//	seq = append(seq[:index], seq[index+1:]...)
//	fmt.Println(seq)
//}
//
//package main
//import "fmt"
//
//func main(){
//	maplit := map[string]int{"one":1,"two":2}
//	var mapAsssigned map[string]float32
//	mapCreated := make(map[string]float32)
//	mapAsssigned = mapCreated
//	mapAsssigned["twp"] = 1
//	maplit["asds"]= 123
//	fmt.Println(maplit)
//	fmt.Println(mapAsssigned)
//	fmt.Println(mapCreated)
//}

//package main
//import "fmt"
//func main(){
//	scene := make(map[string]int)
//	scene["route"] = 66
//	scene["brazil"] = 4
//
//	for k,v := range scene{
//		fmt.Print(k,v)
//	}
//}

//package main
//import "fmt"
//import "sort"
//func main(){
//	scene := make(map[string]int)
//	scene["route"] = 66
//	scene["brazil"] = 4
//	scene["china"] = 960
//	var sceneList []string
//	for k:= range scene{
//		sceneList = append(sceneList,k)
//	}
//	sort.Strings(sceneList)
//	fmt.Println(sceneList)
//}


//package main
//import "fmt"
//func main(){
//	scene := make(map[string]int)
//	scene["route"] = 66
//	scene["brazil"] = 4
//	scene["china"] = 960
//	delete(scene,"route")
//	fmt.Println(scene)
//
//}
//
//package main
//import "fmt"
//import "container/list"
//func main(){
//	l := list.New()
//	l.PushBack("canon")
//	l.PushFront(67)
//	element := l.PushBack("first")
//	l.InsertAfter("high",element)
//	l.InsertBefore("noon", element)
//	fmt.Println(element)
//	fmt.Println(l)
//	l.Remove(element)
//	for i := l.Front();i !=nil; i=i.Next(){
//		fmt.Println(i.Value)
//	}
//}

//package main
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//func main(){
//	inputReader := bufio.NewReader(os.Stdin)
//	fmt.Println("plesase input your name:")
//	input, err := inputReader.ReadString('\n')
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	} else {
//		name := input[:len(input)-2]
//		fmt.Printf("hello what can l do for you %s", name)
//	}
//	for{
//		input ,err := inputReader.ReadString('\n')
//		if err != nil{
//			fmt.Println(err)
//		}
//		input = input[:len(input)-2]
//		input = strings.ToLower(input)
//		switch input {
//		case "":
//				continue
//		case "nothing","bye":
//			fmt.Println("bye!")
//			os.Exit(0)
//		default:
//			fmt.Println("sorry, l did not catch you")
//		}
//	}
//}

//package main
//import "fmt"
//
//type Data struct {
//	complax []int
//	instance InnerData
//	ptr *InnerData
//}
//type InnerData struct {
//	a int
//}
//
//func passByValue(inFunc Data)Data{
//	fmt.Printf("infunc value: %+v\n",inFunc)
//	fmt.Printf("infunc ptr: %p\n", &inFunc)
//	return inFunc
//}
//func main(){
//	//111111
//	var q Data
//	q.complax = []int{1,3,3}
//	//222222
//	dat := new(Data)
//	dat.complax = []int{1,2,3}
//	//333333
//	in := Data{
//		complax: []int{1,2,3},
//		instance:InnerData{a:5},
//		ptr:&InnerData{1},
//	}
//	fmt.Printf("in value %v\n",in)
//	fmt.Printf("in ptr: %p\n", &in)
//	out := passByValue(in)
//	fmt.Printf("out value: %v\n",out)
//	fmt.Printf("out ptr: %p\n", &out)
//}

//package main
//import (
//	"testProject/src/mytest/dohttp"
//)
//func main(){
//	dohttp.Printsome1()
//}
//package main
//import (
//	"fmt"
//	"io/ioutil"
//	"net"
//	"os"
//)
//func main(){
//	fmt.Println(os.Args)
//	if len(os.Args) != 2{
//		fmt.Fprint(os.Stderr, "Usage: %s host:port ", os.Args[0])
//		os.Exit(1)
//	}
//	service := os.Args[1]
//	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
//	checkError(err)
//	conn, err := net.DialTCP("tcp",nil, tcpAddr)
//	checkError(err)
//	//_, err = conn.Write([]byte("123"))
//	//checkError(err)
//	result, err := ioutil.ReadAll(conn)
//	checkError(err)
//	fmt.Println(string(result))
//	os.Exit(0)
//}
//
//func checkError(err error) {
//	if err != nil{
//		fmt.Fprint(os.Stderr, "Fatal error: %s", err.Error())
//		os.Exit(1)
//	}
//}

//
//package main
//import (
//	"encoding/json"
//	"fmt"
//)
//
//func main(){
//	user := make(map[string]int)
//	user["username"] = 1
//	user["password"] = 2
//	jsonStr, _ := json.Marshal(user)
//	fmt.Println(string(jsonStr))
//	str := "{\"in_progress\":0,\"list\":null,\"total\":0,\"waiting\":0}"
//	target := make(map[string]int)
//	json.Unmarshal([]byte(str),&target)
//	fmt.Println(target)
//	fmt.Println(target["list"])
//}
//
//package main
//
//import (
//	"fmt"
//	"net"
//	"regexp"
//	"strings"
//)
//
//func main(){
//	result := make(map[string]bool)
//	netloc := "213.203.221.14:873"
//
//	conn,err := net.Dial("tcp", netloc)
//	if err != nil{
//		return
//	}
//	helloString := ReadTheLine(conn)
//	fmt.Println(helloString)
//	_, err = conn.Write([]byte(helloString))
//	if err != nil{
//		return
//	}
//	_, err = conn.Write([]byte("\n"))
//	if err != nil{
//		return
//	}
//
//
//	for{
//		buf := ReadTheLine(conn)
//		fmt.Println(buf)
//		if strings.Index(string(buf),"@RSYNCD: EXIT") != -1{
//			fmt.Println("no")
//			break
//		} else if strings.Index(string(buf),"@RSYNCD: AUTHRE") != -1{
//			fmt.Println("no")
//			break
//		} else if strings.Index(string(buf), "@ERROR") != -1{
//			fmt.Println("no")
//			break
//		} else {
//			reg := regexp.MustCompile(`[\w]+`)
//			ret := reg.FindAllString(string(buf), 1)
//			fmt.Println(ret)
//			if len(ret) != 0{
//				if CheckAuth(netloc,ret[0]){
//					result[ret[0]] = true
//				} else {
//					result[ret[0]] = false
//				}
//			}
//
//		}
//		fmt.Println(result)
//	}
//	conn.Close()
//}
//
//func ReadTheLine(conn net.Conn) string{
//	tempbuf := ""
//	buf := make([]byte, 1)
//	for{
//		_, err := conn.Read(buf)
//		if err != nil{
//			return ""
//		}
//		tempbuf += string(buf)
//		if string(buf) == string("\n"){
//			break
//		}
//		buf = make([]byte, 1)
//	}
//	return tempbuf
//}
//
//func CheckAuth(loc string, target string) bool{
//	conn,err := net.Dial("tcp", loc)
//	if err != nil{
//		return false
//	}
//	helloString := ReadTheLine(conn)
//	fmt.Println(helloString)
//	_, err = conn.Write([]byte(helloString))
//	if err != nil{
//		return false
//	}
//	_, err = conn.Write([]byte(target + "\n"))
//	if err != nil{
//		return false
//	}
//	for {
//		buf := ReadTheLine(conn)
//		fmt.Println(buf)
//		if strings.Index(string(buf), "@RSYNCD: OK") != -1{
//			return true
//		} else if strings.Index(string(buf),"@RSYNCD: EXIT") != -1{
//			return false
//		} else if strings.Index(string(buf),"@RSYNCD: AUTHRE") != -1{
//			return false
//		} else if strings.Index(string(buf), "@ERROR") != -1{
//			return false
//		} else {
//			return false
//		}
//	}
//}

//package main
//
//import (
//	"encoding/xml"
//	"fmt"
//	"os"
//)
//
//type Website struct {
//	Name string `xml:"name,attr"`
//	Url string
//	Course []string
//}
//func main(){
//	info := Website{"ccccc", "http://11.1.1.",[]string{"asds"}}
//	f, err := os.Create("./info.xml")
//	if err != nil{
//		fmt.Println("file create err!", err.Error())
//		return
//	}
//	defer f.Close()
//	encoder := xml.NewEncoder(f)
//	err = encoder.Encode(info)
//	if err != nil{
//		fmt.Println(err.Error())
//		return
//	} else{
//		fmt.Println("success")
//	}
//}

//package main
//
//import (
//	"encoding/xml"
//	"fmt"
//	"os"
//)
//
//type Website struct {
//	Name string `xml:"name,attr"`
//	Url string
//	Course []string
//}
//
//func main(){
//	file, err := os.Open("./info.xml")
//	if err!= nil{
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//	info := Website{}
//	decoder := xml.NewDecoder(file)
//	err = decoder.Decode(&info)
//	if err != nil{
//		fmt.Println(err)
//		return
//	} else {
//		fmt.Println("success")
//		fmt.Println(info)
//	}
//}
//package main
//
//import (
//	"encoding/gob"
//	"fmt"
//	"os"
//)
//
//func main(){
//	info := map[string]string{
//		"name": "C语言中文网",
//		"website": "ccccc",
//	}
//	name := "demo.gob"
//	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
//	defer File.Close()
//	enc := gob.NewEncoder(File)
//	if err := enc.Encode(info);err != nil{
//		fmt.Println(err)
//	}
//}


//package main
//
//import (
//	"encoding/gob"
//	"fmt"
//	"os"
//)
//
//func main(){
//	var M map[string]string
//	File, _ := os.Open("demo.gob")
//	D := gob.NewDecoder(File)
//	D.Decode(&M)
//	fmt.Println(M)
//}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//func main(){
//	filePath := "./output.txt"
//	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//	str := "assadasdasdasdasdasdsadasdasdasdasdasdasdasdsda"
//	writer := bufio.NewWriter(file)
//	writer.Write([]byte(str))
//	writer.Flush()
//}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main(){
//	filePath := "./output.txt"
//	file, err := os.Open(filePath)
//	defer file.Close()
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	//if file, err := os.Open(filePath);err !=nil{
//	//	fmt.Println(err)
//	//	return
//	//}
//	reader := bufio.NewReader(file)
//	//var content []byte
//	for{
//		byte, err := reader.ReadByte()
//		if err != nil{
//			return
//		}
//		fmt.Print(string(byte))
//	}
//	//fmt.Println(content)
//}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"io"
//	"os"
//)
//
//func main(){
//	file, err := os.Open("./output.txt")
//	if err != nil{
//		fmt.Println(err)
//	}
//	defer file.Close()
//	reader := bufio.NewReader(file)
//	for {
//		str, err := reader.ReadString('\n')
//		if err == io.EOF{
//			break
//		}
//		fmt.Println(str)
//	}
//	fmt.Println("end...")
//}
//
//package main
//
//import "fmt"
//
//func main(){
//	type myint int
//	var a myint
//	a = 1
//	fmt.Println(a)
//}

//
//package main
//
//import (
//	"bytes"
//	"encoding/binary"
//	"fmt"
//	"os"
//)
//
//type Website struct {
//	Url int32
//}
//func main(){
//	file, err := os.Create("output.bin")
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//	for i := 1;i<=10; i++{
//		info := Website{123}
//		var bin_buf bytes.Buffer
//		binary.Write(&bin_buf, binary.LittleEndian, info)
//		fmt.Println(bin_buf.Bytes())
//		b := bin_buf.Bytes()
//		_, err = file.Write(b)
//		if err != nil{
//			fmt.Println(err)
//			return
//		}
//
//	}
//}

//
//package main
//
//import (
//	"bytes"
//	"encoding/binary"
//	"fmt"
//	"os"
//)
//
//type Website struct {
//	Url int32
//}
//
//func readNextBytes(file *os.File, number int) []byte{
//	bytes := make([]byte, number)
//	_, err := file.Read(bytes)
//	if err != nil{
//		fmt.Println(err)
//	}
//	return bytes
//}
//func main(){
//	file, err := os.Open("output.bin")
//	defer file.Close()
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	m := Website{}
//	for i := 1; i<= 10;i++{
//		data := readNextBytes(file, 4)
//		buffer := bytes.NewBuffer(data)
//		err = binary.Read(buffer, binary.LittleEndian,&m)
//		if err != nil{
//			fmt.Println(err)
//		}
//		fmt.Println(m)
//	}
//}

//package main
//
//import (
//	"flag"
//	"fmt"
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"time"
//)
//func main(){
//	flag.Parse()
//	roots := flag.Args()
//	if len(roots) == 0{
//		roots = []string{"."}
//	}
//	//bianli wenjian shu
//	fileSizes := make(chan int64)
//	go func() {
//		for _, root := range roots{
//			walkDir(root, fileSizes)
//		}
//		close(fileSizes)
//	}()
//	//shuchu jieguo
//	fmt.Println("shuchu")
//	var nfiles, nbytes int64
//	for size := range fileSizes{
//		nfiles ++
//		nbytes += size
//	}
//	fmt.Println("shuchu2")
//	printDiskUsage(nfiles,nbytes)
//}
//
//func printDiskUsage(nfiles, nbytes int64){
//	fmt.Printf("%d file, %.1f GB\n", nfiles,float64(nbytes)/1e9)
//
//}
//func walkDir(dir string, fileSizes chan<- int64){
//	time.Sleep(time.Millisecond)
//	for _, entry := range dirents(dir){
//		if entry.IsDir(){
//			subdir := filepath.Join(dir, entry.Name())
//			walkDir(subdir, fileSizes)
//		} else {
//			fileSizes <- entry.Size()
//			fmt.Println(entry.Size(), entry.Name())
//		}
//	}
//}
//
//func dirents(dir string) []os.FileInfo{
//	entries, err := ioutil.ReadDir(dir)
//	if err != nil{
//		fmt.Fprint(os.Stderr,"errerrerr")
//		return nil
//	}
//	return entries
//}

//package main
//import (
//	"flag"
//	"fmt"
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"time"
//)
//
//var verbose = flag.Bool("v", false, "显示详细进度")
//
//func main(){
//	flag.Parse()
//	roots := flag.Args()
//	if len(roots) == 0{
//		roots = []string{"."}
//	}
//	fileSizes := make(chan int64)
//	go func(){
//		for _, root := range roots{
//			walkDir(root, fileSizes)
//		}
//		close(fileSizes)
//	}()
//	var tick <- chan time.Time
//	if *verbose{
//		tick = time.Tick(500* time.Millisecond)
//	}
//	var nfiles,nbytes int64
//	loop:
//		for {
//			select {
//			case size, ok := <- fileSizes:
//				if !ok{
//					break loop
//				}
//				nfiles++
//				nbytes += size
//				case <- tick:
//					printDiskUsage(nfiles,nbytes)
//			}
//		}
//		printDiskUsage(nfiles,nbytes)
//}
//func walkDir(dir string, fileSizes chan<- int64){
//	for _, entry := range dirents(dir){
//		if entry.IsDir(){
//			subdir := filepath.Join(dir, entry.Name())
//			walkDir(subdir, fileSizes)
//		} else {
//			fileSizes <- entry.Size()
//		}
//	}
//}
//func dirents(dir string) []os.FileInfo{
//	entries, err := ioutil.ReadDir(dir)
//	if err != nil{
//		fmt.Println(err)
//		return nil
//	}
//	return entries
//}
//func printDiskUsage(nfiles, nbytes int64){
//	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
//}

//package main
//
//import (
//	"os"
//	"syscall"
//)
//
//type FileLock struct {
//	dir string
//	f *os.File
//}
//func New(dir string) *FileLock{
//	return &FileLock{
//		dir: dir:dir,
//		f:   nil,
//	}
//}
//func (l *FileLock)Lock() error{
//	f, err := os.Open(l.dir)
//	if err != nil{
//		return err
//	}
//	l.f = f
//	err = syscall.Flock(int(f.Fd()),syscall.LOCK_EX|syscall.LOCK_NB)
//
//}

//os包的使用
//os包的使用
//os包的使用
//os包的使用
//package main
//import (
//	"fmt"
//	"os/exec"
//)
//func main(){
//	f, err := exec.LookPath("ls")
//	if err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println(f)
//}

