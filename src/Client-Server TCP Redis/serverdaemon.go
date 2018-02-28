package main
import (
    "fmt"
    "time"
    "log"
    "net"
    "os"
    "os/signal"
    //"strconv"
    "syscall"
    "encoding/binary"

    "github.com/takama/daemon"
    "github.com/mediocregopher/radix.v2/redis"
)
type Service struct {
    daemon.Daemon
}

func (service *Service) Manage() (string, error){
  usage := ""
  interrupt := make(chan os.Signal, 1)
  signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
  listener, err := net.Listen("tcp", ":9977")
  if err!=nil{
     return  "Possibly was a problem with the port binding", err
  }
  //channel for accepted connection
  listen := make(chan net.Conn, 100)

  go acceptConnection(listener, listen)
  for {
        select {
        case conn := <-listen:
            fmt.Println("Calling handleClient")
            go handleClient(conn)
        case killSignal := <-interrupt:
            fmt.Println("Got signal:", killSignal)
            fmt.Println("Stoping listening on ", listener.Addr())
            listener.Close()
            if killSignal == os.Interrupt {
                return "Daemon was interruped by system signal", nil
            }
            return "Daemon was killed", nil
        }
  } 
  return usage, nil

} 

// Accept a client connection and collect it in a channel
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        listen <- conn
    }
}

func handleClient(client net.Conn) {
    for {
        //buf := make([]byte, 1024)
        //numbytes, err := client.Read(buf)
	var passedvalue uint64
        err := binary.Read(client, binary.LittleEndian, &passedvalue)
        //passedvalue := binary.BigEndian.Uint64(buf)
        passedvaluestr := fmt.Sprintf("%v", passedvalue)
        fmt.Println("passedvaluestr",passedvaluestr)
        //if numbytes == 0 || err != nil {
	if err != nil {
            fmt.Println("Returing from client.Read")
            return
        }
        
        conn, err := redis.Dial("tcp", "localhost:6379")
        if err != nil {
           log.Fatal(err)
        }
        fmt.Println("DB connection established")
        defer conn.Close()
        resp := conn.Cmd("HSET", "sysinfo:1", passedvaluestr, passedvaluestr)
        if resp.Err != nil {
           log.Fatal(resp.Err)
        }
	time.Sleep(1 * time.Second)
        fmt.Println("Successfully added entry into database")
    }
}

func main(){
  srv, err := daemon.New("InsertDB", "Database Insertion Daemon")
  if err != nil {
      fmt.Println("Error: ", err)
      os.Exit(1)
  }
  service := &Service{srv}
  status, err :=  service.Manage() 
  if err != nil {
    fmt.Println(status, "\nError: ", err)
    os.Exit(1)
  }
  fmt.Println(status)
}
