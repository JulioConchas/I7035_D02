package main

import "fmt"
import "math/rand"

// julio conchas
// conchasjimenez@gmail.com
// Fibonacci

func main()  {
    var option int
    menu()
    fmt.Scanf("%d",&option)

    switch option {
    case 1: fubonacci()
    case 2: variadic()
    case 3: n_impares_generator()
    case 4: intercambioMenu()
    default: fmt.Println("Practica número:",option," no existe")
    }
}
func menu(){
    fmt.Println("======= Julio Conchas ======")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("========= Practicas ========")
    fmt.Println("[1] Fibonacci.")
    fmt.Println("[2] Variadic Function")
    fmt.Println("[3] Generador # impares")
    fmt.Println("[4] Intercambio")
}
func intercambioMenu(){
    fmt.Println("======== Intercambio =======")
    var a,b int
    fmt.Scanln(&a)
    fmt.Scanln(&b)

    fmt.Println("a:",a,"b:",b)
    intercambio(&a,&b)
    fmt.Println("a:",a,"b:",b)
}
func intercambio(x,y *int){
    var aux int
    aux = *x
    *x = *y
    *y = aux
}
func n_impares_generator(){
    fmt.Println("===== Generador Impares ====")
    nextImpar := generadorImpares()
    fmt.Println(nextImpar())
    fmt.Println(nextImpar())
    fmt.Println(nextImpar())
    fmt.Println(nextImpar())
    fmt.Println(nextImpar())
    fmt.Println(nextImpar())
}
func generadorImpares() func() uint {
    i := uint(1)
    return func () uint{
        var par = i
		i += 2
        return par
    }
}
func variadic(){
    fmt.Println("========== Variadic ========")
    var slice []int

    for i := 0; i < 100;i++ {
        slice = append(slice,rand.Intn(50))
    }
    // fmt.Println(slice)
    // fmt.Println("==========================")
    fmt.Println(bigest_number(slice...))
}
func bigest_number(n ...int) int{
    var b int
    var flag = true
    b = 0
    for i := 0;i < len(n) ;i++ {
        for j := 0; j < len(n); j++ {
            if n[i] < n[j] {
                flag = false
            }
        }
        if flag {
            b = n[i]
        } else {
            flag = true
        }
    }
    return b
}
func fubonacci(){
    fmt.Println("========= Fibonacci ========")
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
