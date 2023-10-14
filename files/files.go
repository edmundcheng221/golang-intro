package files

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read() {
	dat, err := os.ReadFile("../read.txt")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Create("../write.txt")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered string\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
