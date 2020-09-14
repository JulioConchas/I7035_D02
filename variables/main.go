package main

import "fmt"

func main(){
    var option int
    menu()
    fmt.Scanf("%d",&option)

    switch option {
    case 1: fmt.Println("Area del cuadrado")
    case 2: fmt.Println("Area triangulo")
    case 3: fmt.Println("Area circulo")
    case 4: fmt.Println("Fahrenheit to Celcius")
    default: fmt.Println("Practica número:",option," no existe")
    }
}

func menu(){
    fmt.Println("======= Julio Conchas ======")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("========= Practicas ========")
    fmt.Println("[1] Area del cuadrado")
    fmt.Println("[2] Area del triángulo")
    fmt.Println("[3] Area del círculo")
    fmt.Println("[4] Fahrenheit a Celcius")
}
