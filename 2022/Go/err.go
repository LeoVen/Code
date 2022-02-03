package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can't take negative square root: %s", fmt.Sprint(float64(e)))
}

func Sqrt(x float64) (res float64, err error) {
	if x < 0.0 {
		err = ErrNegativeSqrt(x)
		return
	}
	z := float64(1.0)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
