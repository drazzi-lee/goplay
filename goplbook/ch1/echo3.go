package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sep := "\n"
	fmt.Println(strings.Join(os.Args[1:], sep))
}
