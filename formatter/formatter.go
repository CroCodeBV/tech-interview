package formatter

import "fmt"

const (
    Reset   = "\033[0m"
    Red     = "\033[31m"
    Green   = "\033[32m"
    Yellow  = "\033[33m"
    Blue    = "\033[34m"
    Magenta = "\033[35m"
    Cyan    = "\033[36m"
)

func Format(op string, in, out int) string {
    return fmt.Sprintf("%s%s%s of %s%d%s = %s%d%s",
        Blue, op, Reset,
        Yellow, in, Reset,
        Green, out, Reset)
}
