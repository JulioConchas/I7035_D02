package main

import (
    "fmt"
    "./multimedia"
    "os"
)
// Author: Julio Conchas
// Main  : conchasjimenez@gmail.com
func main(){
    var option int
    md := multimedia.ContenidoWeb{}

    for  true {
        menu()
        fmt.Scanf("%d",&option)
        switch option {
        case 1: agregar_imagen(&md)
        case 2: agregar_audio(&md)
        case 3: agregar_video(&md)
        case 4: show(&md)
        case 5: os.Exit(1)
        default: fmt.Println("Practica número:",option," no existe")
        }
    }
}
func menu(){
    fmt.Println("======= Julio Conchas ======")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("========== GO VI ===========")
    fmt.Println("[1] Agregar Imagen.")
    fmt.Println("[2] Agregar Audio.")
    fmt.Println("[3] Agregar Video.")
    fmt.Println("[4] Mostrar Contenido Web.")
    fmt.Println("[5] Salir")
}
func show(m *multimedia.ContenidoWeb){
    m.Mostrar()
}
func agregar_imagen(m *multimedia.ContenidoWeb){
    var t,f,c string
    fmt.Println("Ingrese Titulo, formato, canales: ")
    fmt.Scanln(&t)
    fmt.Scanln(&f)
    fmt.Scanln(&c)
    im := multimedia.Imagen{Titulo:t,Formato:f,Canales:c}

    m.Multimedia = append(m.Multimedia,&im)
}
func agregar_audio(m *multimedia.ContenidoWeb){
    var t,f string
    var d int
    fmt.Println("Ingrese Titulo, formato, duración: ")
    fmt.Scanln(&t)
    fmt.Scanln(&f)
    fmt.Scanln(&d)
    au := multimedia.Audio{Titulo:t,Formato:f,Duracion:d}

    m.Multimedia = append(m.Multimedia,&au)
}
func agregar_video(m *multimedia.ContenidoWeb){
    var t,f string
    var fr int
    fmt.Println("Ingrese Titulo, formato, frames: ")
    fmt.Scanln(&t)
    fmt.Scanln(&f)
    fmt.Scanln(&fr)
    vi := multimedia.Video{Titulo:t,Formato:f,Frames:fr}

    m.Multimedia = append(m.Multimedia,&vi)
}
