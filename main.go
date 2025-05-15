package main

import (
	"fmt"
	"os"

	"tech-interview/calculator"
	"tech-interview/formatter"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: orders <number>")
        os.Exit(1)
    }

    n, err := calculator.ParseInt(os.Args[1])
    if err != nil {
        fmt.Println("Invalid number:", os.Args[1])
        os.Exit(1)
    }

    sq := calculator.Square(n)
    out := formatter.Format("square", n, sq)
    fmt.Println(out)
}
