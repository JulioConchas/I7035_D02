package main

import (
    "fmt"
    "net/rpc"
    "bufio"
    "os"
    "time"
)

//Author: Julio Conchas
//Mail  : conchasjimenez@gmail.com

func cliente(nikname string){
    var op int64
    var sms string
    var result int
    var result_str string

    c, err := rpc.Dial("tcp",":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    err = c.Call("Server.Connection",nikname,&result)
    if err != nil {
        fmt.Println(err)
    } else {
        if result != 1 {
            fmt.Println("Error al conectar al server")
        }
    }
    for {

        menu()
        fmt.Scanln(&op)
        switch op {
        case 1:
            fmt.Println("Mensaje: ")
            scanner := bufio.NewScanner(os.Stdin)
            scanner.Scan() // use `for scanner.Scan()` to keep reading
            sms = scanner.Text()
            err = c.Call("Server.Send",nikname+","+sms,&result)
            if err != nil {
                fmt.Println(err)
            } else {
                if result != 1 {
                    fmt.Println("Error al enviar el mensaje")
                }
            }
        case 2:
            fmt.Println("Enviando archivo")
        case 3:
            fmt.Println("Mostrando chat")
                err = c.Call("Server.Print_chat",nikname,&result_str)
                if err != nil {
                    fmt.Println(err)
                } else {
                    fmt.Println(result_str)
                }
        case 0:
            fmt.Println("Saliendo")
            err = c.Call("Server.DesConectar",nikname,&result)
            if err != nil {
                fmt.Println(err)
            } else {
                if result != 1 {
                    fmt.Println("Error al conectar al server")
                }
            }
            return
        }
    }
    c.Close()
}
func main(){
    var nikname string
    fmt.Print("Nickname: ")
    fmt.Scanln(&nikname)
    go print_msg(nikname)
    cliente(nikname)
}
func print_msg(nikname string){
    var sms string
    var result_str string

    for {
        c, err := rpc.Dial("tcp",":9999")
        if err != nil {
            fmt.Println(err)
            return
        }
        err = c.Call("Server.Print_chat",nikname,&result_str)
        if err != nil {
            fmt.Println(err)
        } else {
            if result_str != "_trash" {
                if sms != result_str {
                    fmt.Println(result_str)
                    sms = result_str
                }
            }
        }
        c.Close()
        time.Sleep(time.Millisecond * 500)
    }
}
func menu() {
    fmt.Println("1) Enviar Mensaje")
    fmt.Println("2) Enviar Archivo")
    fmt.Println("3) Mostar chat")
    fmt.Println("0) Exit")
}
