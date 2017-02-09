package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    a := make([][]uint8, dy)
    for x := range a {
        b := make([]uint8, dx)
        for y:= range b {
            b[y] = uint8(2*x)
        }
        a[x] = b
    }
    return a
}

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Println("Hello, 世界")
	fmt.Printf("Hello, %s, %d\n", "Drazzi", 291)
	pic.Show(Pic)
}
