package main

import "fmt"

func main() {
    var a int
    as := "hello"
    b := 1
    fmt.Printf("%s: %d   %s, %d", "helloworld", a, as, b)

    const (
        en = iota
        bb
        c
        d
    )
    fmt.Printf("(%d) 1 val", en)
    fmt.Printf("(%d) 2 val", bb)
    fmt.Printf("(%d) 3 val", c)
    fmt.Printf("(%d) 4 val", d)

    for i:=0; i< 11; i++ {
        if i>5 {
            break
        }
        fmt.Println(i)
    }

    list := []string { " hello" , "dudue"}

    for k, v := range list {
        fmt.Print(k, v)
    }
    fmt.Print(list)


    switch {
        case en == 0:
             fmt.Printf("Enum  0")
            //  fallthrough
        case en == 1:
            fmt.Printf("Enum 1")
    }

    var one_d [3]int;
    //  = {1,2,3}
    // one_d := [...]int{1,2,3}
    // var two_d [2][2]int = { {1,2}, {3,4}}

    fmt.Println("nONe D = ", one_d)
    // fmt.Println("Two D = ", two_d)
    // Slice , growing array
    // Reference types are created with make

    fmt.Println("Slices...")

    var array [100]int
    slice := array[0:99]

    slice[9] = 99
    fmt.Println(slice)

    s1 := append(slice, 1,2,3,4)
    s2 := append(slice, s1...)

    fmt.Printf("Size of s1 = %d, cap = %d", len(s1), cap(s1))
    fmt.Printf("Size of s2 = %d, cap = %d", len(s2), cap(s2))

    // Last line ends with comma
    // Map is reference type
    monthdays := map[string]int {
        "Jan" : 31,
        "Feb" : 28,
    }

    fmt.Println(monthdays)

}


