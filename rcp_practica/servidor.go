package main

import (
	"fmt"
    "strconv"
    "strings"
	"net"
	"net/rpc"
    // "./server"
)

type Server struct {

}

var Materias map[string]map[string]float64
var Alumno map[string]float64

func (s *Server) Calificacion_materia(name_maria_data string,reply *string) error {
    split := strings.Split(name_maria_data,",")
    elFloat , err := strconv.ParseFloat(split[2], 64)
    if err == nil {
        Alumno[split[0]] = elFloat
        Materias[split[1]] = Alumno
    }
    *reply = "Alumno: "+split[0]+" Materia: "+split[1]+" data: "+split[2]
    return nil
}
func (s *Server) Print (str string, reply *string) error {
    for key,val := range Materias{
        str1 := fmt.Sprintf("%s=\"%s\"", key, val)
        str = str + str1
    }
    *reply = str
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

func main (){
    Alumno = make(map[string]float64)
    Materias = make(map[string]map[string]float64)
    
    go servidor()

	var input string
	fmt.Scanln(&input)
}
