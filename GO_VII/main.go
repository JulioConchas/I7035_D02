package main

import "fmt"

// Author: Julio Conchas
// Main  : conchasjimenez@gmail.com

func main(){
    s := []string{}
    var aux string

    fmt.Println("Ingresa el String..")
    fmt.Scanln(&aux)
    s = append(s,aux)
    for true {
        fmt.Println("Otro string? [y/n]")
        fmt.Scanln(&aux)
        if aux=="y"||aux=="Y" {
            fmt.Println("Ingresele pues...")
            fmt.Scanln(&aux)
            s = append(s,aux)
        } else {
            break
        }
    }
    fmt.Println(s)
}
