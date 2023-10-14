package main

import (
	"EdmundsBankai/golang-intro/files"
	helloworld "EdmundsBankai/golang-intro/hello-world"
	http_test "EdmundsBankai/golang-intro/http"
	num "EdmundsBankai/golang-intro/numberParsing"
	"EdmundsBankai/golang-intro/random"
	t "EdmundsBankai/golang-intro/time"
	"EdmundsBankai/golang-intro/url"
	"cmp"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"slices"
	s "strings"
	"time"
)

var print = fmt.Println

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type person struct {
	name string
	age  int
}

type rect struct {
	width  int
	height int
}

type geometry interface {
	getArea() int
}

func routine1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "hdsbfds")
	}
}
func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out)) // string() takes a slice of bytes or runes and converts it into a string type
	slice := make([]string, 3)
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println(slice, len(slice))
	slice = append(slice, "d")
	fmt.Println(slice, len(slice), "new slice")
	copied := make([]string, len(slice))
	copy(copied, slice)
	fmt.Println(copied, "copied slice")
	fmt.Println(slices.Equal(copied, slice))
	maps()
	rangeFunction()
	variadic(1, 2, 3, 5)
	closures()
	pointers()
	bob := person{"Bob", 20}
	fmt.Printf("%s is %v years old\n", bob.name, bob.age)
	//   compute area
	rectange := rect{width: 10, height: 9}
	area := rectange.getArea()
	fmt.Println(area, "is the area")
	// interfaces
	r := rect{width: 3, height: 4}
	measure(&r)
	fmt.Println(errorHandling(42))
	fmt.Println(errorHandling(43))

	go routine1()
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
	bufferedChannels()
	timeouts()
	channelLooping()
	sorting()
	// catch and recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	stringOperations()
	encoding()
	helloworld.Helloworld()
	currTime := t.GetTime()
	print(currTime.Year())
	print(currTime.Month())
	print(currTime.Day())
	print(currTime.Hour())
	print(currTime.Location())
	t.GetEpoch()
	t.FormatTime()
	random.GetRandomNumbers()
	num.Parse()
	url.FormatUrl()
	files.Read()
	http_test.Http()
	// spawning commands
	dateCmd := exec.Command("date")
	dateOut, dateErr := dateCmd.Output()
	if dateErr != nil {
		panic("error executing date command")
	}
	fmt.Println(string(dateOut))
	// os.Exit(3) // exits with status 3, the defers will not execute

}

func maps() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	// getting a value that doesn't exist returns a 0
	v3 := m["k3"]
	if v3 != 0 {
		fmt.Println("v3:", v3)
	} else {
		fmt.Println("value not found")
	}
	delete(m, "k2")
	fmt.Println(m)
	_, prs := m["k2"] // prs = if key exists
	fmt.Println("prs:", prs)
}

func rangeFunction() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println(sum)
	}
	// iterating over a map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func variadic(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closures() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())
}

func updateVal(val int) {
	val = 20
}

func updatePtr(val *int) {
	*val = 20
}

func pointers() {
	example := 10
	fmt.Println(example)
	updateVal(example)
	fmt.Println(example) // unchanged bc we pass by value
	updatePtr(&example)
	fmt.Println(example) // changed to 20 because we pass by reference
}

// methods
func (r *rect) getArea() int {
	return r.height * r.width
}

// interfaces

func measure(g geometry) {
	fmt.Println(g.getArea())
}

func errorHandling(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cannot work with 42")
	}
	return 20, nil
}

func bufferedChannels() {
	msgs := make(chan string, 2)
	msgs <- "test1"
	msgs <- "test2"
	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
}

// we can specify the channel direction for better type safety
// func ping(pings chan<- string, msg string) { // here we specify that the ping channel only accepts a channel for sending values
// 	pings <- msg
// }

// func pong(pings <-chan string, pongs chan<- string) {
// 	msg := <-pings
// 	pongs <- msg
// }

// select allows for awaiting multiple go routines

func timeouts() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	default: // implements non-blocking sends and receives
		fmt.Println("no message sent")
	}
	close(c1) // close channel to prevent sends

}

func channelLooping() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}

}

// we can also use wait groups to wait for multiple go routines to finish

func sorting() {
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)
	chars := []string{"c", "a", "b"}
	slices.Sort(chars)
	fmt.Println("Characters:   ", chars)
	// custom sorting
	fruits := []string{"peach", "banana", "kiwi"}
	lengths := func(char1, char2 string) int {
		return cmp.Compare(len(char2), len(char1)) // reverse order based on length of the word
	}
	slices.SortFunc(fruits, lengths)
	fmt.Println(fruits)
	_, err := os.Create("file.go")
	os.Remove("file.go")
	if err != nil {
		panic(err) // similar to throw error
	}

}

// defer - equivalent to the finally in js
func deference() {
	// f := createFile("/tmp/defer.txt")
	// defer closeFile(f) // this is executed at the end of the function
	// writeFile(f)
}

func mayPanic() {
	if false {
		panic("some error")
	}
}

func stringOperations() {
	print("Contains:  ", s.Contains("test", "es"))
	print("Count:     ", s.Count("test", "t"))
	print("HasPrefix: ", s.HasPrefix("test", "te"))
	print("HasSuffix: ", s.HasSuffix("test", "st"))
	print("Index:     ", s.Index("test", "e"))
	print("Join:      ", s.Join([]string{"a", "b"}, "-"))
	print("Repeat:    ", s.Repeat("a", 5))
	print("Replace:   ", s.Replace("foo", "o", "0", -1))
	print("Replace:   ", s.Replace("foo", "o", "0", 1))
	print("Split:     ", s.Split("a-b-c-d-e", "-"))
	print("ToLower:   ", s.ToLower("TEST"))
	print("ToUpper:   ", s.ToUpper("test"))
}

func encoding() {
	// conversion of data between json and go data structures
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B), reflect.TypeOf(res2B))
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res, reflect.TypeOf(res))
}
