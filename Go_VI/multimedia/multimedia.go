package multimedia

import ( "fmt" )

type  ContenidoWeb struct{
    Multimedia []Multimedia_
}
type Multimedia_ interface{
    Mostrar()
}
type Imagen struct{
    Titulo string
    Formato string
    Canales string
}
type Audio struct{
    Titulo string
    Formato string
    Duracion int
}
type Video struct{
    Titulo string
    Formato string
    Frames int
}
// methodos
func (i *Imagen) Mostrar() {
    fmt.Println(i.Titulo+" "+i.Formato+" "+i.Canales)
}
func (a *Audio) Mostrar() {
    fmt.Println(a.Titulo+" "+a.Formato+" ",a.Duracion)
}
func (v *Video) Mostrar() {
    fmt.Println(v.Titulo+" "+v.Formato+" ",v.Frames)
}
func (cw *ContenidoWeb) Mostrar() {
    for _, m := range cw.Multimedia{
        m.Mostrar()
    }
}
