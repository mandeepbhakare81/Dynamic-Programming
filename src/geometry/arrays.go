package main
import "fmt"

func main(){
  a:=[...] int{17,18,19}
  fmt.Println(a)
  b:=[3]string{"USA","India","UK"}
  c:=b
  c[1]="Singapore"
  fmt.Println(b)
  fmt.Println(c)
  sum:=float64(0)
  d:=[...]float64{10.5,12.3,34.2,5.2,1.5}
  for i,v:=range d{
	fmt.Println(d[i])
        sum=sum+v
  }
  fmt.Println("Sum of all members of array d =",sum)
}
