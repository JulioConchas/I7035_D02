package main

import (
    "fmt"
    "net"
    "encoding/gob"
    "./proceso"
    "time"
)

var p1 = proceso.Proceso{0,0}
var p2 = proceso.Proceso{1,0}
var p3 = proceso.Proceso{2,0}
var p4 = proceso.Proceso{3,0}
var p5 = proceso.Proceso{4,0}

var clientes = 0
var stop_p1 = 0
var stop_p2 = 0
var stop_p3 = 0
var stop_p4 = 0
var stop_p5 = 0

func servidor(){
    s,err := net.Listen("tcp",":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        go the_proces()
        c, err := s.Accept()  //escucha peticiones
        if err != nil {
            fmt.Println(err)
            continue
        }
        go handleClient(c)
    }
}
func handleClient(c net.Conn){
    var coming_proces proceso.Proceso

    switch clientes {
        case 0:
            stop_p1=1
            err := gob.NewEncoder(c).Encode(p1)
            if err != nil{
                fmt.Println(err)
            }
        case 1:
            stop_p2=1
            err := gob.NewEncoder(c).Encode(p2)
            if err != nil{
                fmt.Println(err)
            }
        case 2:
            stop_p3=1
            err := gob.NewEncoder(c).Encode(p3)
            if err != nil{
                fmt.Println(err)
            }
        case 3:
            stop_p4=1
            err := gob.NewEncoder(c).Encode(p4)
            if err != nil{
                fmt.Println(err)
            }
        case 4:
            stop_p5=1
            err := gob.NewEncoder(c).Encode(p5)
            if err != nil{
                fmt.Println(err)
            }
    }
    clientes=clientes+1
    err := gob.NewDecoder(c).Decode(&coming_proces)
    if err != nil{
        fmt.Println(err)
    } else{
        clientes=clientes-1
    }
    restart_Proces(coming_proces)
}
func main(){
    go servidor()
    var input string
    fmt.Scanln(&input)

}
func the_proces(){
    for{
        fmt.Println("--------------------------")
        if stop_p1 != 1 {
            p1.Increase()
        }
        if stop_p2 != 1 {
            p2.Increase()
        }
        if stop_p3 != 1 {
            p3.Increase()
        }
        if stop_p4 != 1 {
            p4.Increase()
        }
        if stop_p5 != 1 {
            p5.Increase()
        }
        time.Sleep(time.Millisecond * 500)
    }
}
func restart_Proces(p proceso.Proceso)  {
    if p.GetID() == 0 {
        fmt.Println("",p.GetI())
        p1.SetI(p.GetI())
        stop_p1 = 0
    } else if p.GetID() == 1 {
        p2.SetI(p.GetI())
        stop_p2 = 0
    } else if p.GetID() == 2 {
        p3.SetI(p.GetI())
        stop_p3 = 0
    }else if p.GetID() == 3 {
        p4.SetI(p.GetI())
        stop_p4 = 0
    } else if p.GetID() == 4 {
        p5.SetI(p.GetI())
        stop_p5 = 0
    }
}
