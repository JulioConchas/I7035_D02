package main

import (
    "fmt"
    "net"
    "encoding/gob"
)
type Persona struct {
    Nombre string
    Email []string
}

func cliente(persona Persona){
    c,err := net.Dial("tcp",":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    err = gob.NewEncoder(c).Encode(persona)
    if err != nil{
        fmt.Println(err)
    }
    c.Close()
}
func main(){
    persona := Persona{
        Nombre : "Julio Conchas",
        Email: []string{
            "conchasjimenez@gmail.com",
            "alien.287@hotmail.com",
        },
    }
    go cliente(persona)

    var input string
    fmt.Scanln(&input)
}
