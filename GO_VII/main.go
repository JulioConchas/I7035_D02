package main

import (
    "fmt"
    "os"
    "strings"
    "sort"
)

// Author: Julio Conchas
// Main  : conchasjimenez@gmail.com

func main(){
    s := []string{}
    var aux string

    file_ascendente,err := os.Create("ascendente.txt")
    if err != nil{
        fmt.Println(err)
        return
    }
    defer file_ascendente.Close()

    file_descendente,err := os.Create("descendente.txt")
    if err != nil{
        fmt.Println(err)
        return
    }
    defer file_descendente.Close()

    fmt.Println("Ingresa el String..")
    fmt.Scanln(&aux)
    s = append(s,aux)
    for true {
        fmt.Println("Otro string? [y/n]")
        fmt.Scanln(&aux)
        if aux=="y"||aux=="Y" {
            fmt.Println("Ingresele pues...")
            fmt.Scanln(&aux)
            s = append(s,aux)
        } else {
            break
        }
    }
    sort.Sort(ByUpward(s))
    file_ascendente.WriteString(strings.Join(s,"\n"))
    sort.Sort(ByFalling(s))
    file_descendente.WriteString(strings.Join(s,"\n"))
}

type ByUpward []string
func (a ByUpward) Len() int { return len(a) }
func (a ByUpward) Swap(i,j int) { a[i], a[j] = a[j], a[i] }
func (a ByUpward) Less(i,j int) bool { return a[i] < a[j] }

type ByFalling []string
func (a ByFalling) Len() int { return len(a) }
func (a ByFalling) Swap(i,j int) { a[i], a[j] = a[j], a[i] }
func (a ByFalling) Less(i,j int) bool { return a[i] > a[j] }
