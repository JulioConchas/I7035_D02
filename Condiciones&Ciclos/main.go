package main

import "fmt"

func main(){
    var option int
    menu()
    fmt.Scanf("%d",&option)

    switch option {
    case 1: prueba_clase()
    // case 2: area_triangulo()
    // case 3: area_circulo()
    // case 4: fahrenheit_to_celcius()
    default: fmt.Println("Practica número:",option," no existe")
    }
}
func menu(){
    fmt.Println("======= Julio Conchas ======")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("========= Practicas ========")
    fmt.Println("[1] Prueba de clase")
}
func prueba_clase(){
    var temp int

    fmt.Print("Temp: ")
    fmt.Scan(&temp)

    if temp < 0 {
        fmt.Println("Está helado")
    } else if temp >= 0 && temp < 12 {
        fmt.Println("Está haciendo frío")
    } else if temp >= 12 && temp < 18 {
        fmt.Println("Está agusto")
    }else {
        fmt.Println("Está caluroso")
    }
}
