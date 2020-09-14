package main

import "fmt"

func main(){
    var option int
    menu()
    fmt.Scanf("%d",&option)

    switch option {
    case 1: area_cuadrado()
    case 2: area_triangulo()
    case 3: area_circulo()
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
func area_cuadrado(){
    var lado,resultado float32

    fmt.Println("===== Area del Cuadrado ====")
    fmt.Println("Ingresa el lado: ")
    fmt.Scanf("%g",&lado)
    resultado = lado * lado
    fmt.Println("Area = ",resultado,"cm")
}
func area_triangulo(){
    var base,altura,area float32

    fmt.Println("==== Area del Triángulo ====")
    fmt.Println("Ingresa la base: ")
    fmt.Scanf("%g",&base)
    fmt.Println("Ingresa la altura: ")
    fmt.Scanf("%g",&altura)
    area = ( base * altura ) / 2
    fmt.Println("Area = ",area,"cm")
}
func area_circulo(){
    var radio,area,pi float32
    pi = 3.14

    fmt.Println("===== Area del Circulo =====")
    fmt.Println("Ingresa el radio: ")
    fmt.Scanf("%g",&radio)
    area = pi * (radio*radio)
    fmt.Println("Area = ",area,"cm")
}
