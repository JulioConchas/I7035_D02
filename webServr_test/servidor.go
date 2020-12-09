package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main(){
    http.HandleFunc("/",root)
    fmt.Println("Arrancando el servidor...")
    http.ListenAndServe(":9000",nil)
}


//funciones
func root(res http.ResponseWriter, req *http.Request){
    res.Header().Set(
        "Content-Type",
        "text/html",
    )
    fmt.Fprintf(
        res,
        cargarHtml("./html/index.html"),
    )
}
func cargarHtml(a string) string{
    html, _ := ioutil.ReadFile(a)
    return string(html)
}
