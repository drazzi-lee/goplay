package main

import (
    "bufio"
    "database/sql"
    "fmt"
    "os"
    "strings"
    _ "github.com/go-sql-driver/mysql"

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

type Person struct {
    Name string
    Age  int
}

// what's the difference between *Person and Person?
// if the param p defined as *Person, the *Person type can call the func String()
// but the Person type can NOT call the func String()
func (p *Person) String() string {
    return fmt.Sprintf("BB __%v (%v years)", p.Name, p.Age)
}

type Hello struct{}

/*
func (h Hello) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}*/

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
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

	a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}

	//&z call by referrence, so it can call func String() as type *Person 
    fmt.Println(a.String(), &z)

	// goroutine and channel
	nums := []int{7, 2, 8, -9, 4, 0, 19, 72, 12, 28}
	c := make(chan int)
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)
	x, y := <- c, <- c
	fmt.Println(x, y, x+y)

	// range and close for channel
	ch := make(chan int, 20) // define with channel length 10
	fmt.Println(ch)
	go fibonacci2(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}

	// http
	/*
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}*/

    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }

    db, err := sql.Open("mysql", "root:1234@mysql:127.0.0.1")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

}
