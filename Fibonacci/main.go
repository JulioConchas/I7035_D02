package main

import "fmt"

// julio conchas
// conchasjimenez@gmail.com
// Fibonacci

func main()  {
    var option int
    menu()
    fmt.Scanf("%d",&option)

    switch option {
    case 1: fibonacci()
    default: fmt.Println("Practica número:",option," no existe")
    }
}
func menu(){
    fmt.Println("======= Julio Conchas ======")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("========= Practicas ========")
    fmt.Println("[1] Fibonacci.")
    fmt.Println("[2] Número e.")
}
