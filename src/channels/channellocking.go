package main
import (
  "fmt"
  "sync"
)
var x int = 0

func increment(w *sync.WaitGroup, ch chan bool){
  ch<-true
   x=x+1
   <-ch
   w.Done()
}

func main(){
  var w sync.WaitGroup
  ch:=make(chan bool, 1)
  for i:=0; i<1000; i++ {
   w.Add(1)
   increment(&w, ch)
  }
  w.Wait()
  fmt.Println("The value of x",x)
}
    
