package main
import (
   "fmt"
)
func find(num int, nums...int){
 fmt.Println("Numnber to find is",num,"from the numbers",nums)
 found:=false
 for i,v:=range nums{
     if v==num{
      fmt.Println(num,"found at index",i)
      found=true
      break
     }
 }
 if !found{
  fmt.Println(num,"is not present in",nums)
 }
}

func change(s...string){
  s[0] = "Go"
  s = append(s, "playground")
  fmt.Println(s);
}

func main(){
  find(8,4,3,5,6,7);
  nums:=[]int{2,7,5,12,20,25,78}
  find(25,nums...)

  welcome := []string{"hello", "world"}
  fmt.Println(welcome)
  change(welcome...)
  fmt.Println(welcome) 
} 
