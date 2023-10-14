package num

import (
	"fmt"
	"strconv"
)

func Parse() {
	// parse functions return error on bad input
	f, _ := strconv.ParseFloat("1.234", 64) // we want 64 bits of precision
	fmt.Println(f)
	i, _ := strconv.ParseInt("123", 0, 64) // int requires 3 args, 0 is saying that we want it to infer the base
	fmt.Println(i)
}
