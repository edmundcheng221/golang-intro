package random

import (
	"fmt"
	"math/rand"
)

var print = fmt.Println

func GetRandomNumbers() {
	print(rand.Intn(100))
	print(rand.Float64())
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	print(r2)
}
