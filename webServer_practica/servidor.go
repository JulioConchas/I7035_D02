package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "net/http"
)


//Author: Julio Conchas
//Mail  : conchasjimenez@gmail.com

//STRUCTS
type Alumno struct{
    Name string
    Materias map[string]float64
}
func (a *Alumno) AddMateria(name string,data float64){
    a.Materias[name] = data
}

//variables globales
var losAlumnos []Alumno
var m map[string]float64

func main(){

    m = make(map[string]float64)

    http.HandleFunc("/",root)
    http.HandleFunc("/add_form", add_form)
    http.HandleFunc("/response", response)
    http.HandleFunc("/promedio_alumno", promedio_alumno)
    http.HandleFunc("/promedio_general", promedio_general)
    http.HandleFunc("/promedio_materia", promedio_materia)
    fmt.Println("Arrancando el servidor...")
    http.ListenAndServe(":9000",nil)
}


//funciones
func cargarHtml(a string) string{
    html, _ := ioutil.ReadFile(a)
    return string(html)
}
func isAlumno(a string) int{
    for i := 0;i < len(losAlumnos);i++ {
        if losAlumnos[i].Name == a {
            return i
        }
    }
    return -1
}
func promedio_x_alumno(name string) string {
    var suma float64
    var nMaterias float64
    var promedio float64

    index := isAlumno(name)
    if index != -1 {
        for key,val := range losAlumnos[index].Materias {
            fmt.Println(name+":"+key)
            suma = suma + val
            nMaterias = nMaterias + 1
        }
    }
    promedio = suma / nMaterias
    elString := fmt.Sprintf("%f", promedio)
    return elString
}
func promedio_x_general() string {
    var suma float64
    var nMaterias float64
    var suma_gral float64
    var promedio float64

    for i := 0; i < len(losAlumnos); i++ {
        for key,val := range losAlumnos[i].Materias {
            fmt.Println(":"+key)
            suma = suma + val
            nMaterias = nMaterias + 1
        }
        suma_gral=suma_gral+(suma/nMaterias)
        suma=0
        nMaterias=0
    }
    promedio = suma_gral / float64(len(losAlumnos))
    elString := fmt.Sprintf("%f", promedio)
    return elString
}
func promedio_x_materia(materia string) string{
    var suma float64
    var nMaterias float64
    var promedio float64

    for i := 0; i < len(losAlumnos); i++ {
        if _,aux := losAlumnos[i].Materias[materia]; aux {
            suma = suma + losAlumnos[i].Materias[materia]
            nMaterias = nMaterias + 1
        }
    }
    promedio = suma / nMaterias
    elString := fmt.Sprintf("%f", promedio)
    return elString
}
//paginas

func root(res http.ResponseWriter, req *http.Request){
    res.Header().Set(
        "Content-Type",
        "text/html",
    )
    fmt.Fprintf(
        res,
        cargarHtml("./web/index.html"),
    )
}
func add_form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("./web/add_form.html"),
	)
}
func response(res http.ResponseWriter, req *http.Request){
    fmt.Println(req.Method)
    switch req.Method {
    case "POST":
        if req.FormValue("calificacion") != "" {
            fmt.Println("Agregando alumno")
            if err := req.ParseForm(); err != nil {
    			fmt.Fprintf(res, "ParseForm() error %v", err)
    			return
    		}
    		fmt.Println(req.PostForm)
            elFloat , err := strconv.ParseFloat(req.FormValue("calificacion"), 64)
            if err == nil {
                index := isAlumno(req.FormValue("nombre"))
                if index != -1 {
                    losAlumnos[index].AddMateria(req.FormValue("materia"),elFloat)
                } else{
                    m = make(map[string]float64)
                    m[req.FormValue("materia")] = elFloat
                    a := Alumno{req.FormValue("nombre"),m}
                    losAlumnos = append(losAlumnos,a)
                }
            }
            fmt.Println(losAlumnos)
            res.Header().Set(
                "Content-Type",
                "text/html",
            )
            fmt.Fprintf(
    			res,
    			cargarHtml("./web/response.html"),
    			req.FormValue("nombre"),
                req.FormValue("materia"),
                req.FormValue("calificacion"),
    		)
        } else if req.FormValue("name") != "" {
            fmt.Println("Promedio por alumno")
            fmt.Println(req.PostForm)
            res.Header().Set(
                "Content-Type",
                "text/html",
            )
            fmt.Fprintf(
    			res,
    			cargarHtml("./web/response_promedio.html"),
                req.FormValue("name"),
    			promedio_x_alumno(req.FormValue("name")),
    		)
        } else if req.FormValue("subject") != "" {
            fmt.Println("Promedio por Materia")
            fmt.Println(req.PostForm)
            res.Header().Set(
                "Content-Type",
                "text/html",
            )
            fmt.Fprintf(
    			res,
    			cargarHtml("./web/response_materia.html"),
                req.FormValue("subject"),
    			promedio_x_materia(req.FormValue("subject")),
    		)
        }
    }
}
func promedio_alumno(res http.ResponseWriter, req *http.Request){
    res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("./web/promedio_alumno_form.html"),
	)
}
func promedio_general(res http.ResponseWriter, req *http.Request)  {
    fmt.Println(req.Method)
    switch req.Method {
    case "GET":
        res.Header().Set(
    		"Content-Type",
    		"text/html",
    	)
    	fmt.Fprintf(
    		res,
    		cargarHtml("./web/promedio_general.html"),
            promedio_x_general(),
    	)
    }
}
func promedio_materia(res http.ResponseWriter, req *http.Request)  {
    res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("./web/promedio_materia_form.html"),
	)
}
