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
    case 1: fubonacci()
    default: fmt.Println("Practica número:",option," no existe")
    }
}
func menu(){
    fmt.Println("======= Julio Conchas ======")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("========= Practicas ========")
    fmt.Println("[1] Fibonacci.")
}
func fubonacci(){
    var x int
    fmt.Scanln(&x)
    fmt.Println(fibonacci_recursivo(x))
}
func fibonacci_recursivo(n int) int{
    if n < 2 {
        return n
    } else {
        return (fibonacci_recursivo(n-1) + fibonacci_recursivo(n-2))
    }
}
