package main
import(
  "fmt"
)


func calcSquares(num int, sqrsm chan int){
  sum:=0
  var c int 
  for num != 0 {
    c=num%10
    num = num/10
    sum = sum+c*c
  }
  sqrsm<-sum
} 

func calcCube(num int, cubesm chan int){
 sum:=0
 var c int 
 for num != 0 {
  c=num%10
  num = num/10
  sum=sum+c*c*c
 }
 cubesm<-sum
}

func main(){
  number:=589
  sqrsm:= make(chan int)
  cubesm:= make( chan int)
  go calcSquares(number, sqrsm)
  go calcCube(number, cubesm)
  squaresum:=<-sqrsm
  cubesum:=<-cubesm
  TotalSum:=squaresum+cubesum
  fmt.Println("Total sum=",TotalSum)
}
