package main

// Author: Julio Conchas
// Main  : conchasjimenez@gmail.com

import (
    "fmt"
    "os"
)

func main(){
    var opt int
    for {
        menu()
        fmt.Scanf("%d",&opt)
        switch opt {
        case 1: fmt.Println("agregando new job")
        case 2: fmt.Println("mostrando los jobs")
        case 3: fmt.Println("eliminando job")
        case 4: os.Exit(1)
        }
    }
}
func menu(){
    fmt.Println("============================")
    fmt.Println("=      Julio Conchas       =")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("=        Gorutines         =")
    fmt.Println("============================")
    fmt.Println("[1] Add Job.")
    fmt.Println("[2] Display Jobs.")
    fmt.Println("[3] Finish Jobs.")
    fmt.Println("[4] Salir")
}
