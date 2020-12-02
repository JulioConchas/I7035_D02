package proceso

// Author: Julio Conchas
// Main  : conchasjimenez@gmail.com

import (
    "fmt"
    "time"
    )

type Proceso struct {
    Id uint64
    I uint64
}

func (p *Proceso) Increase(c chan int){
    for {
        p.I = p.I + 1
        msg := <- c
        if msg == -1 {
            fmt.Printf( "  %d        %d\n", p.Id, p.I)
        }
        m := uint64(msg)
        if m == p.Id {
            return
        }
        time.Sleep(time.Millisecond * 500)
    }
}
