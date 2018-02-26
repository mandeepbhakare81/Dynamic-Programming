package main
import (
    "fmt"
    "github.com/mediocregopher/radix.v2/redis"
    //"github.com/garyburd/redigo/redis" 
    //"errors"
    //"time"
    "log"
)
type SData struct{
   stype string
   svalue float64  
}
func InsertIntoDB(dataCh chan SData){
  conn, err := redis.Dial("tcp", "localhost:6379")
  if err != nil {
     log.Fatal(err)
  }
  SData<-dataCh
  defer conn.Close()
  var sensortype string
  var value float64
  resp := conn.Cmd("HSET", SData.stype, SData.value) 
  if resp.Err != nil {
    log.Fatal(resp.Err)
  }
  defer conn.Close()

}

func TempSensor(){
    tmpch=make(chan SData, 1)
    InsertIntoDB(tmpch)
  
}

func MotionSensor(){
    motch=make(chan SData, 1)
    InsertIntoDB(motch) 
}

func NoiseSensor(){
     noisech=make(chan SData, 1)
     InsertIntoDB(noisech)
}

func main(){
    // the Str() helper method to convert the reply to a string.
    title, err := conn.Cmd("HGET", "album:1", "title").Str()
    if err != nil {
        log.Fatal(err)
    }
    sd:= SData{temperature, 29.4}
    DataCh=make(chan SData,1)
    DataCh<-sd
    go TempSensor(DataCh)
    go MotionSensor(DataCh)
    go NoiseSensor(DataCh)
    
    fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)
}
