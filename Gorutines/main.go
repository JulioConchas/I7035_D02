package main

// Author: Julio Conchas
// Main  : conchasjimenez@gmail.com

import (
    "fmt"
    "os"
    "./proceso"
)

var id_count = uint64(0);
var ch_flag = -2

func main(){
    var opt int
    var c chan int = make(chan int)

    for {
        menu()
        fmt.Scanf("%d",&opt)
        switch opt {
        case 1: add_job(c)
        case 2: display_jobs(c)
        case 3: delete_jobs(c)
        case 4: os.Exit(1)
        default: fmt.Println(opt, " this option does not exists")
        }
    }
}
func menu(){
    fmt.Println("============================")
    fmt.Println("=      Julio Conchas       =")
    fmt.Println("= conchasjimenez@gmail.com =")
    fmt.Println("=        Gorutines         =")
    fmt.Println("============================")
    fmt.Println("[1] Add Job.")
    fmt.Println("[2] Display Jobs.")
    fmt.Println("[3] Finish Jobs.")
    fmt.Println("[4] Salir")
}
func add_job(c chan int){
    fmt.Println(id_count," : job added successfully!")
    p := proceso.Proceso{id_count,0}
    go p.Increase(c)
    id_count = id_count + 1

}
func display_jobs(c chan int){
    var opt string
    fmt.Println("Abailable Jobs")
    fmt.Println("============================")
    fmt.Println("  ID      Value ")
    ch_flag = -1
    go courier(c)

    fmt.Scanln(&opt)
    if opt == "x" {
        ch_flag = -2
    }
}
func delete_jobs(c chan int){
    var op int
    fmt.Println("Job id? ")
    fmt.Scanf("%d",&op)
    cast := uint64(op)
    if cast > id_count {
        fmt.Println("Job does not exists")
    } else {
        ch_flag = op
    }
}
func courier(c chan int){
    for {
        c <- ch_flag
    }
}
