package formatter

import (
	"fmt"
	"strings"
)

func Cube(n int) int {
	return n*n
}

const (
    Reset   = "\033[0m"
    C    = "\033[36m"
    M = "\033[35m"
    Y  = "\033[33m"
    R     = "\033[31m"
    G   = "\033[32m"
    B    = "\033[34m"
)

func Format(op string, in, out int) string {
    return fmt.Sprintf("%s%s%s of %s%d%s = %s%d%s",
        B, op, Reset,
        Y, in, Reset,
        G, out, Reset)
}

func Trim(s string) string {
  return strings.TrimSpace(s)
}
