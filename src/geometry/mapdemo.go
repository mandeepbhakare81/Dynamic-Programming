package main
import (
  "fmt"
)

func main(){
   personSalary := map[string]int {
    "Mandeep": 50000,
    "Ganesh":90000,
  }
    //personSalary = make(map[string]int)
    personSalary["Ramesh"]  = 70000
    personSalary["Suresh"]  = 80000
    fmt.Println("map personSalary is created", personSalary)
    value, ok := personSalary["Mandeep"]
    if ok == true{
    	employee:="Mandep"
    	fmt.Println("Salary of the",employee,"is",value)
    } else {
        fmt.Println("Employee Mandep is not present in the personSalary")
    }
}
