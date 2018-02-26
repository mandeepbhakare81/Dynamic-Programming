package main
import (
    "fmt"
    "github.com/mediocregopher/radix.v2/redis"
    //"github.com/garyburd/redigo/redis" 
    //"errors"
    //"time"
    "log"
    "sync"
)
func InsertIntoDB(dataCh chan map[string]string, wg *sync.WaitGroup){
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
  fmt.Println("Verifying the data inserted inot the databse")
  reply, err := conn.Cmd("HGETALL", "SysInfo:1").Map()
  fmt.Println("err",err)
  if err == nil {
    fmt.Println("Data successfully fetched from Database")
    for first, second := range reply {
       fmt.Println("Key:",first,"Value:",second)
    }
  }
  wg.Done()
}



func InitSystem(systemch chan int, mapch chan map[string]string, wg *sync.WaitGroup){
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
  wg.Done() 
}

func RunProgram(systemch chan int, mapch chan map[string]string, wg *sync.WaitGroup){
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
    wg.Done()
}

func ShutDownSystem(systemch chan int, mapch chan map[string]string, wg *sync.WaitGroup){
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
  wg.Done()  
}

func main() {
    var wg sync.WaitGroup                 //Creating a waitgroup
    SystemMap := make(map[string]string)  //Created map to store system data
    fmt.Println("Created SystemMap")
    mapch := make(chan map[string]string)  //created channel to send map data between different function calls
    fmt.Println("Created mapch")
    go func(){
      mapch<-SystemMap            //Insert map into the channel
    }()
    fmt.Println("Inserted SystemMap in mapch")
    
    systemch:= make(chan int, 1)    //Created channel to synchronize among functions. It's buffer length is kept one so that only one thread can read or write into it at a time
    fmt.Println("Created systemch")
    go func(){
      systemch<-1        
    }()
    fmt.Println("Calling goroutines")
    wg.Add(1)                            //Incrementing waitgroup counter 
    go InitSystem(systemch, mapch, &wg)
    fmt.Println("Out of InitSystem")

    wg.Add(1)
    go RunProgram(systemch, mapch, &wg)
    fmt.Println("Out of  RunProgram")

    wg.Add(1)
    go ShutDownSystem(systemch, mapch, &wg)
    fmt.Println("Out of ShutDownSystem")
    for key,value := range SystemMap {   //Iterating through data added by different goroutines
       fmt.Println("key=",key,"value",value)
    }

    fmt.Println("Inserting map values into database")
    wg.Add(1)
    go InsertIntoDB(mapch, &wg)
    fmt.Println("Out of InsertIntoDB")
    wg.Wait()
    fmt.Println("Program completed")
}
