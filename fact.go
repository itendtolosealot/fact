package fact

import (
	"errors"
	"fmt"
)

func Factorial(n int) (error, int) {
	if n < 0 {
		fmt.Printf("Validation Error. Number %d < 0\n", n)
		return errors.New("Number not positive"), -1
	} else if n == 0 {
		return nil, 1
	} else if n > 25 {
		fmt.Printf("Validation Error. Number %d > 25\n", n)
		return errors.New("Number too high"), -1
	} else {
		fact := 1
		for i := 1; i <= n; i++ {
			fact = fact * i
		}
		return nil, fact
	}
}
