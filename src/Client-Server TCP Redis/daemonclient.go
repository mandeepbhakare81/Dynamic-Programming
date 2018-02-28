package main

import (
    "fmt"
    "log"
    "net"
    "sync"
    "time"
    //"strconv"
    "encoding/binary"
    //"os"
    //"os/signal"
    //"syscall"
    //"github.com/mediocregopher/radix.v2/redis"
)

func SendToDaemon(intch chan int, wg *sync.WaitGroup, m *sync.Mutex){
  conn, err := net.Dial("tcp", "127.0.0.1:9977")
  if err != nil {
     log.Fatal(err)
  }
  fmt.Println("Daemon  connection established")
  for {
   select {
     case v:=<-intch:
      v64:=int64(v)
      //conn.Write([]byte(strconv.Itoa(v)))
	err := binary.Write(conn, binary.LittleEndian, v64)
	if err != nil {
           log.Fatal(err)
        }
     default:
      time.Sleep(1 * time.Second)
      fmt.Println("Executing default case")
    }
  }
  wg.Done()
}

func goroutine1(onech chan  int, wg *sync.WaitGroup, m *sync.Mutex){
      go func(){
       for i:=0; i < 100; i++{
        m.Lock()
        onech<-i
	m.Unlock()
        time.Sleep(1 * time.Second)
       }
      }()
   wg.Done()
}

func goroutine2(twoch chan  int, wg *sync.WaitGroup, m *sync.Mutex){
      go func(){
       for i:=200; i < 300; i++{
        m.Lock()
        twoch<-i
	m.Unlock()
        time.Sleep(2 * time.Second)
       }
      }()
  wg.Done()
}


func main() {
    var wg sync.WaitGroup                 //Creating a waitgroup
    var m sync.Mutex
    datach := make(chan int, 1)
    fmt.Println("Created mapch")

    fmt.Println("Calling goroutines")
    wg.Add(1)                            //Incrementing waitgroup counter
    go goroutine1(datach, &wg, &m)
    fmt.Println("Out of InitSystem")

    wg.Add(1)
    go goroutine2(datach, &wg, &m)
    fmt.Println("Out of  RunProgram")

    fmt.Println("Inserting map values into database")
    wg.Add(1)
    go SendToDaemon(datach, &wg, &m)

    wg.Wait()
    fmt.Println("Program completed")
}
  
