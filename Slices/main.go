package main

import "fmt"

// julio conchas
// conchasjimenez@gmail.com
// test_burbuja

func main()  {
    s := []int64{1, -10, 90, 14, -100, -2}
    Burbuja(s)
}
func Burbuja(s []int64){
    for i := 0;i < len(s);i++ {
        fmt.Println(s[i])
    }
}
