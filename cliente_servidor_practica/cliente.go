package main

import (
    "fmt"
    "net"
    "encoding/gob"
    "./proceso"
    "time"
    "os/signal"
    "os"
    "syscall"
)

func cliente(){
    var proceso proceso.Proceso
    ch := make(chan os.Signal, 1)
    signal.Notify(ch,syscall.SIGTERM,syscall.SIGINT);

    c,err := net.Dial("tcp",":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    err = gob.NewDecoder(c).Decode(&proceso)
    if err != nil{
        fmt.Println(err)
    } else{
        //go runProcess(proceso)
        go func () {
            for{
                go proceso.Increase()
                time.Sleep(time.Millisecond * 500)
            }
        }()
    }
    // looking of an end signal "command + c"
    for sig := range ch{
        switch sig {
        case syscall.SIGINT:
            //fmt.Println("",proceso)
            err := gob.NewEncoder(c).Encode(proceso)
            if err != nil{
                fmt.Println(err)
            }
            fmt.Println("terminando presicillo1")
            os.Exit(0)
        }
    }
    c.Close()
}
func main(){
    go cliente()

    var input string
    fmt.Scanln(&input)
}
// func runProcess(p proceso.Proceso){
//
// }
