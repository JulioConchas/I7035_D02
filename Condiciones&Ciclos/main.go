package main

import "fmt"

func main(){
    var option int
    menu()
    fmt.Scanf("%d",&option)

    switch option {
    case 1: signo_zodiacal()
    case 2: numero_e()
    default: fmt.Println("Practica número:",option," no existe")
    }
}
func menu(){
    fmt.Println("======= Julio Conchas ======")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("========= Practicas ========")
    fmt.Println("[1] Signo zodiacal.")
    fmt.Println("[2] Número e.")
}
func signo_zodiacal(){
    var dia int
    var mes int

    fmt.Scanf("%d",&dia)
    fmt.Scanf("%d",&mes)

    if dia <= 31 && mes <= 12 {
        if ( dia >= 21 && mes == 3 )||( dia <= 20 && mes == 4 ) {
            fmt.Println("Aries")
        } else if ( dia >= 21 && mes == 4 )||( dia <= 20 && mes == 5 ) {
            fmt.Println("Tauro")
        } else if ( dia >= 21 && mes == 5 )||( dia <= 21 && mes == 6 ) {
            fmt.Println("Geminis")
        } else if ( dia >= 22 && mes == 6 )||( dia <= 22 && mes == 7 ) {
            fmt.Println("Cancer")
        } else if ( dia >= 23 && mes == 7 )||( dia <= 22 && mes == 8 ) {
            fmt.Println("Leo")
        } else if ( dia >= 23 && mes == 8 )||( dia <= 22 && mes == 9 ) {
            fmt.Println("Viergo")
        } else if ( dia >= 23 && mes == 9 )||( dia <= 22 && mes == 10 ) {
            fmt.Println("Libra")
        } else if ( dia >= 23 && mes == 10 )||( dia <= 22 && mes == 11 ) {
            fmt.Println("Escorpio")
        }  else if ( dia >= 23 && mes == 11 )||( dia <= 21 && mes == 12 ) {
            fmt.Println("Sagitario")
        } else if ( dia >= 22 && mes == 12 )||( dia <= 20 && mes == 1 ) {
            fmt.Println("Capricornio")
        } else if ( dia >= 21 && mes == 1 )||( dia <= 18 && mes == 2 ) {
            fmt.Println("Acuario")
        } else if ( dia >= 19 && mes == 2 )||( dia <= 20 && mes == 3 ) {
            fmt.Println("Piscis")
        }
    } else {
        fmt.Println(" No existe mes con: días ",dia," meses",mes," lol ")
    }

}
func numero_e(){
    var e float32
    e = 0
    for i := 0; i < 30; i++ {
        e = e + 1/float32(factorial(i));
    }
    fmt.Println("e = ",e)
}
func factorial(n int) int{
    var f int;
    var i int;

    if n == 0{
        f = 1;
    } else{
        f = 1
        for i = n; i >= 1; i-- {
            f = f * i;
        }
    }
    return f
}
