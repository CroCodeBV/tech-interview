package main

import (
	"fmt"
	"os"

	"flag"
	"tech-interview/calculator"
	"tech-interview/formatter"
)

func main() {
	cube := flag.Bool("cube", false, "print cube")
	flag.Parse()

    if len(os.Args) < 2 {
        fmt.Println("Usage: orders <number>")
        os.Exit(1)
    }

    n, err := calculator.ParseInt(os.Args[1])
    if err != nil {
        fmt.Println("Invalid number:", os.Args[1])
        os.Exit(1)
    }
  fmt.Println(*cube) // TODO: remove before committing!
    if *cube {
    cu := formatter.Cube(n)
    out := formatter.Format("cube", n, cu)
    fmt.Println(out)
    } else {
    sq := calculator.Square(n)
    out := formatter.Format("square", n, sq)
    fmt.Println(out)
    }
}
