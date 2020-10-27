package main

import "fmt"
// Author: Julio Conchas
// Main  : conchasjimenez@gmail.com

type Imagen struct{
    titulo string
    formato string
    canales string
}
type Audio struct{
    titulo string
    formato string
    duracion int
}
type Video struct{
    titulo string
    formato string
    frames int
}
// methodos
func (i *Imagen) mostrar {
    fmt.Println(i.titulo)
    fmt.Println(i.formato)
    fmt.Println(i.canales)
}
func (a *Audio) mostrar {
    fmt.Println(a.titulo)
    fmt.Println(a.formato)
    fmt.Println(i.duracion)
}
func (v *Video) mostrar {
    fmt.Println(v.titulo)
    fmt.Println(v.formato)
    fmt.Println(v.frames)
}
func main(){

}
