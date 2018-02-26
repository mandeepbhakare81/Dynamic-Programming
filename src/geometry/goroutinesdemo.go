package main
import (
  "fmt"
  "time"
)
func hello(a chan int){
  fmt.Println("Inside hello")
  fmt.Println("hello goroutine is going into sleep")
  time.Sleep(4 * time.Second)
  fmt.Println("hello goroutine has awakened from sleep")
  data:=<-a
  fmt.Println("Data is read from a")
  fmt.Println("data=",data)
}

func main(){
  var a chan int
  if a == nil{
    fmt.Println("Channel is nil")
    a = make(chan int)
    fmt.Printf("Type of a is %T", a)
    fmt.Println()
  }
  fmt.Println("Inside main")
  go hello(a)
  fmt.Println("Back to main")
  a <- 1
  //time.Sleep(1 * time.Second)
  fmt.Println("Inside main")
}
