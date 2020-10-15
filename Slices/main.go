package main

import "fmt"

// julio conchas
// conchasjimenez@gmail.com
// test_burbuja

func main()  {
    s := []int64{1, -10, 90, 14, -100, -2}
    Burbuja(s)
    fmt.Println("main: ",s)
}
func Burbuja(s []int64){
    var aux int64
    for i := 0;i < len(s) ;i++ {
        for j := 0;j < (len(s) - i - 1);j++ {
            if s[j] > s[j+1] {
                aux = s[j]
                s[j] = s[j+1]
                s[j+1] = aux
            }
        }
    }
}
