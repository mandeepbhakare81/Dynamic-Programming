package main
import (
    "fmt"
    "github.com/mediocregopher/radix.v2/redis"
    //"github.com/garyburd/redigo/redis" 
    //"errors"
    "time"
    "log"
)
type SData struct{
   stype string
   svalue string  
}
func InsertIntoDB(dataCh chan map[string]string){
  conn, err := redis.Dial("tcp", "localhost:6379")
  if err != nil {
     log.Fatal(err)
  }
  fmt.Println("DB connection established")
  defer conn.Close()
  SystemMap:=<-dataCh
  fmt.Println("Data read from dataCh")
  for first, second := range SystemMap {
      fmt.Println("first",first,"second",second)
      resp := conn.Cmd("HMSET","SysInfo:1",first, second)
      if resp.Err != nil {
         log.Fatal(resp.Err)
      }
  } 
  fmt.Println("Successfully added map entries into database")
}

func InitSystem(systemch chan int, mapch chan map[string]string){
      fmt.Println("Inside InitSystem")
  if <-systemch==1 {
      fmt.Println("Initializing the system here")
      SystemMap:=<-mapch
      fmt.Println("Read from mapch")
      SystemMap["OS"] = "Solaris"
      SystemMap["Processor"] = "Sparc"
      fmt.Println("Assigned values to SystemMap")
      go func(){
        mapch<-SystemMap
      }()
      go func(){
        systemch<-2
      }()
      fmt.Println("Exiting InitSystem")
  }
   
}

func RunProgram(systemch chan int, mapch chan map[string]string){
    fmt.Println("Inside RunProgram")
    if <-systemch==2 {
      fmt.Println("Running the program here")
      SystemMap:=<-mapch
      fmt.Println("Read from mapch")
      SystemMap["Status"] = "Running"
      SystemMap["Program"] = "VEN"
      fmt.Println("Assigned values to SystemMap")
      go func(){
        mapch<-SystemMap
      }()
      go func(){
        systemch<-3
      }()
      fmt.Println("Exiting RunProgram")
    }
}

func ShutDownSystem(systemch chan int, mapch chan map[string]string){
  if <-systemch==3{
      fmt.Println("Shutting down the system here")
      SystemMap:=<-mapch
      SystemMap["Error"] = "nil"
      SystemMap["Log"] = "/var/log"
      go func(){
        mapch<-SystemMap
      }()
      go func(){
        systemch<-3
      }()
      fmt.Println("Exiting ShutDownSystem")
  }
    
}

func main() {
    // the Str() helper method to convert the reply to a string.
    //title, err := conn.Cmd("HGET", "album:1", "title").Str()
    //if err != nil {
      //  log.Fatal(err)
    //}
    c := make(chan int)
    go func(){
      c <- 42
    }()
    val := <-c
    fmt.Println("Staring in main",val)
    SystemMap := make(map[string]string)
    fmt.Println("Created SystemMap")
    mapch := make(chan map[string]string)
    fmt.Println("Created mapch")
    go func(){
      mapch<-SystemMap
    }()
    fmt.Println("Inserted SystemMap in mapch")
    
    systemch:= make(chan int, 1)
    fmt.Println("Created systemch")
    go func(){
      systemch<-1
    }()
    //done := make(chan bool)
    fmt.Println("Calling goroutines")

    go InitSystem(systemch, mapch)
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("Out of InitSystem")

    go RunProgram(systemch, mapch)
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("Out of  RunProgram")

    go ShutDownSystem(systemch, mapch)
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("Out of ShutDownSystem")
    for key,value := range SystemMap {
       fmt.Println("key=",key,"value",value)
    }

    fmt.Println("Inserting into database")
    InsertIntoDB(mapch)
    fmt.Println("Out of InsertIntoDB")
    //fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)
}
