package main

import (
	"fmt"
    "strconv"
    "strings"
	"net"
	"net/rpc"
    // "./server"
)
var alumnos []Alumno
var m map[string]float64

type Server struct {

}
type Alumno struct{
    Name string
    Materias map[string]float64
}
func (a *Alumno) AddMateria(name string,data float64){
    a.Materias[name] = data
}

func (s *Server) Calificacion_materia(name_maria_data string,reply *string) error {
    split := strings.Split(name_maria_data,",")
    elFloat , err := strconv.ParseFloat(split[2], 64)
    if err == nil {
        index := isAlumno(split[0])
        if index != -1 {
            alumnos[index].AddMateria(split[1],elFloat)
        } else{
            m = make(map[string]float64)
            m[split[1]] = elFloat
            a := Alumno{split[0],m}
            alumnos = append(alumnos,a)
        }
    }
    *reply = "Alumno: "+split[0]+" Materia: "+split[1]+" data: "+split[2]
    return nil
}
func (s *Server) Promedio_alumno(name string,reply *string) error {
    // var suma float64
    // var nMaterias float64
    // var promedio float64
    // nMaterias = 0
    // for key,val := range Materias{
    //     fmt.Println(key)
    //     for alumno, calificacion := range val {
    //         if alumno == name {
    //             suma = suma + calificacion
    //         }
    //     }
    //     nMaterias = nMaterias + 1
    // }
    // promedio = suma / nMaterias
    // elString := fmt.Sprintf("%f", promedio)
    // *reply = elString
    *reply ="promedio"
    return nil
}
func (s *Server) Promedio_general(str string,reply *string) error {
    // var suma float64
    // var nAlumno float64
    // var promedio float64
    // nAlumno = 0
    // for key,val := range Materias{
    //     fmt.Println(key)
    //     for alumno, calificacion := range val {
    //         fmt.Println(alumno)
    //         suma = suma + calificacion
    //         nAlumno = nAlumno + 1
    //     }
    // }
    // promedio = suma / nAlumno
    // elString := fmt.Sprintf("%f", promedio)
    // *reply = elString
    *reply ="promedio"
    return nil
}
func (s *Server) Promedio_materia(name string,reply *string) error {
    // var suma float64
    // var nAlumno float64
    // var promedio float64
    // nAlumno = 0
    // for key,val := range Materias{
    //     fmt.Println(key)
    //     if key == name {
    //         for alumno, calificacion := range val {
    //             fmt.Println(alumno)
    //             suma = suma + calificacion
    //             nAlumno = nAlumno + 1
    //         }
    //     }
    // }
    // promedio = suma / nAlumno
    // elString := fmt.Sprintf("%f", promedio)
    // *reply = elString
    *reply ="promedio"
    return nil
}
func (s *Server) Print (str string, reply *string) error {
    // for key,val := range Materias{
    //     str1 := fmt.Sprintf("%s=\"%s\"", key, val)
    //     str = str + str1
    // }
    str = fmt.Sprintf("%s", alumnos)
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
func isAlumno(a string) int{
    for i := 0;i < len(alumnos);i++ {
        if alumnos[i].Name == a {
            return i
        }
    }
    return -1
}
func main (){
    m = make(map[string]float64)
    //Materias = make(map[string]map[string]float64)

    go servidor()

	var input string
	fmt.Scanln(&input)
}
