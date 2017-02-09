package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
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

func WordCount(s string) map[string]int {
    s_arr := strings.Fields(s)
    s_map := make(map[string]int)
    for i := 0; i < len(s_arr); i++ {
        if s_map[s_arr[i]] == 0 {
            s_map[s_arr[i]] = 1
        } else {
            s_map[s_arr[i]]++
        }
    }
    return s_map
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		i, j = j, i + j
		return i
	}
}

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Println("Hello, 世界")
	fmt.Printf("Hello, %s, %d\n", "Drazzi", 291)
	pic.Show(Pic)

	wc.Test(WordCount)

	//closure
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		//fmt.Println(fibonacci())
		fmt.Println(f())
	}
}
