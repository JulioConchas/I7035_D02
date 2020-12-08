package main

import (
	"fmt"
	"net/rpc"
)

func cliente(){
    c, err := rpc.Dial("tcp",":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    var op int64
    for {
        menu()
        fmt.Scanln(&op)
        switch op {
        case 1:
            var name string
            var materia string
            var calificacion string
            var result string

            fmt.Println("Name: ")
            fmt.Scanln(&name)
            fmt.Println("Materia: ")
            fmt.Scanln(&materia)
            fmt.Println("Calificacion: ")
            fmt.Scanln(&calificacion)

            err = c.Call("Server.Calificacion_materia",name+","+materia+","+calificacion,&result)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Registro: ",result)
            }
        case 2:
            var name string
            var result string

            fmt.Println("Name: ")
            fmt.Scanln(&name)

            err = c.Call("Server.Promedio_alumno",name,&result)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Promedio=",result)
            }
        case 3:
            var result string

            err = c.Call("Server.Promedio_general"," ",&result)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Promedio=",result)
            }
        case 4:
            var name_materia string
            var result string

            err = c.Call("Server.Promedio_materia",name_materia,&result)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Promedio=",result)
            }
        case 5:
            var result string
            err = c.Call("Server.Print","->",&result)
            if err != nil {
                fmt.Println(err)
            }else {
                fmt.Println("Materias=",result)
            }
        case 0:
            return
        }
    }
}
func main(){
    cliente()
}
func menu() {
    fmt.Println("1) Agregar calificaciones")
    fmt.Println("2) Promedio alumno")
    fmt.Println("3) Promedio alumnos")
    fmt.Println("4) Promedio materia")
    fmt.Println("5) Mostrar materias")

    fmt.Println("0) Exit")
}
