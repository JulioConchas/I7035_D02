package main

import (
    "fmt"
    "net"
    "net/rpc"
    "strings"
)

//Author: Julio Conchas
//Mail  : conchasjimenez@gmail.com

type Server struct { }

type Mensaje struct {
    From string
    Sms string
}

var message_mgr []Mensaje

func (s *Server) Connection(name string,reply *int) error {
    fmt.Println("Se conectó: " + name)
    *reply = 1
    return nil
}
func (s *Server) Send(name_sms string,reply *int) error {
    split := strings.Split(name_sms,",")
    m := Mensaje{split[0],split[1]}
    message_mgr = append(message_mgr,m)
    //fmt.Println(message_mgr)
    *reply = 1
    return nil
}
func (s *Server) Print_chat(name string,reply *string) error {
    var sms string
    *reply = "_trash"
    for i := 0; i < len(message_mgr); i++ {
        if name == message_mgr[i].From  {
            sms = sms + "tu : " + message_mgr[i].Sms + "\n"
        } else{
            sms = sms + message_mgr[i].From + " envió : " + message_mgr[i].Sms + "\n"
        }
    }
    *reply = sms
    return nil
}
func (s *Server) DesConectar(name string,reply *int) error {
    fmt.Println("Se desconecto: " + name)
    *reply = 1
    return nil
}
func servidor(){
    rpc.Register(new(Server))
    ln , err := net.Listen("tcp",":9999")
    if err != nil{
        fmt.Println(err)
    }
    for{
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        go rpc.ServeConn(c)
    }
}

func main(){
    go servidor()

    var input string
    fmt.Scanln(&input)

}
