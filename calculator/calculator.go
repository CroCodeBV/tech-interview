package calculator

import (
	"errors"
	"strconv"
)

func ParseInt(s string) (int, error) {
    i, err := strconv.Atoi(s)
    if err != nil {
        return 0, errors.New("not a number")
    }
    return i, nil
}

func Square(x int) int {
    return x * x
}
