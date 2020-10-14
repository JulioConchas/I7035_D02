package main

import "fmt"

func main(){
    var option int
    menu()
    fmt.Scanf("%d",&option)

    switch option {
    case 1: prueba_clase()
    case 2: area_triangulo()
    case 3: area_circulo()
    case 4: fahrenheit_to_celcius()
    default: fmt.Println("Practica número:",option," no existe")
    }
}
func menu(){
    fmt.Println("======= Julio Conchas ======")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("========= Practicas ========")
    fmt.Println("[1] Prueba de clase")
    fmt.Println("[2] Area del triángulo")
    fmt.Println("[3] Area del círculo")
    fmt.Println("[4] Fahrenheit a Celsius")
}
func prueba_clase(){
    
}
