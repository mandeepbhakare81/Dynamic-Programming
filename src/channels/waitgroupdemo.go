package main
import (
   "fmt"
   "sync"
   //"time"
)

func Process( i int, wg *sync.WaitGroup){
  fmt.Println("Executing process", i)
  wg.Done()
}

  
func main(){
  fmt.Println()
  no:=3
  var wg sync.WaitGroup
  for i:=0; i<no; i++ {
    wg.Add(1)
    go Process( i, &wg)
  }
    wg.Wait()
  fmt.Println("ALL three routines have finished processing")
}
