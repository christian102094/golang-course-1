package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"sync"
)

func print(x interface{}) {
	fmt.Printf("%v, %T\n", x, x)
}

func block1() {
	var i = 42
	var i2 int8 = -123
	var j string = strconv.Itoa(i)
	var b bool = true
	b2 := 1 == 2
	d := 3e3
	s := "abcde"
	bs := []byte(s)
	r := 'a'
	const (
		io = iota + 5
		io2
	)

	print(i)
	print(i2)
	print(j)
	print(b)
	print(b2)
	print(d)
	print(bs)
	print(r)
	print(io)
	print(io2)
}

func block2() {
	a := make([]int, 4) // b[1] = 3
	b := []int{4, 5, 6, 7, 8}
	a = append(a, b...)

	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Cap: %v\n", cap(a))
	fmt.Printf("Cap: %v\n", a)
	// fmt.Println("Grades: %v", b)
}

func block3() {
	states := make(map[string]int)
	states = map[string]int{
		"a1": 1,
		"a2": 3,
	}
	states["a2"] = 2
	states["a3"] = 33

	delete(states, "a1")

	pop, ok := states["nonexisting"]

	fmt.Println(states)
	fmt.Println(states["a1"])
	fmt.Println(states["nonexisting"])
	fmt.Println(pop, ok)
	fmt.Println(len(states))
}

func block4() {
	type Doctor struct {
		number     int
		name       string
		companions []string
	}

	aDoctor := Doctor{
		number:     3,
		name:       "Jon",
		companions: []string{"a", "b", "c"},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.name)
	fmt.Println(aDoctor.companions[1])
	fmt.Println(struct{ name string }{name: "Donson"})
	fmt.Println(&struct{ name string }{name: "Donson"})
}

func block5() {
	// Interfaces: behavior, use interchangeable
	// Embedding: reuse

	type Animal struct {
		Name   string `required max: "100"`
		Origin string
	}

	type Bird struct {
		Animal
		Speed  float32
		CanFly bool
	}

	b := Bird{}
	// b.Name = "Emy"
	b.Origin = "Australia"
	b.Speed = 458
	b.CanFly = false

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")

	fmt.Println(b)
	fmt.Println(b.Name, b.Speed)
	fmt.Println(field.Tag)
}

func block6() {
	state := map[string]int{
		"a1": 1,
		"a2": 2,
	}
	if pop, ok := state["a1"]; ok || false {
		fmt.Println(pop)
		fmt.Println(state)
	}

	s1 := "little"
	s2 := "little"

	if s1 == s2 {
		fmt.Println("equal strings")
	}

	switch i := 2 + 3; i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5:
		fmt.Println("3 or 4 or 5")
	default:
		fmt.Println("none")
	}

	j := 10
	switch {
	case j <= 10:
		fmt.Println("<= 10")
		// fallthrough this is logicless
	case j <= 20:
		fmt.Println("<= 20")
	default:
		fmt.Println("")
	}

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("int")
		break
		fmt.Println("int2")
	case float64:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other type")
	}
}

func block7() {
Loop:
	for i, j := 1, 1; i < 5; i++ {
		// continue
		fmt.Println(i, j)
		break Loop
	}

	s := []string{"a", "b", "c"}

	for k, v := range s {
		fmt.Println(k, v)

	}

	states := map[string]int{"a": 11, "b": 22, "c": 33}

	for k, v := range states {
		fmt.Println(k, v)
	}
}

func block8() {
	// defer
	// fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")

	// defer 2
	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()

	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// panic 1
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in f", err)
			// panic(err)
		}
	}()
	panic("panicked")
	fmt.Println("end")

	// panic 2
	// simple server listener
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err)
	// }
}

func block9() {
	// var a int = 42
	// var b *int = &a
	// fmt.Println(a, b)
	// a = 27
	// fmt.Println(&a, *b)
	// *b = 14
	// fmt.Println(a, *b)

	type myStruct struct {
		foo int
	}

	var ms *myStruct = new(myStruct)
	(*ms).foo = 42
	fmt.Println(ms.foo) // syntactic sugar

	// map and slices are pointers to underlying data
}
func sum(values ...int) *int {
	total := 0
	for _, value := range values {
		total += value
	}
	return &total
}

func divideFloat(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

// Method
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func block10() {
	fmt.Println(*sum(1, 2, 3, 4, 5))
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(*sum(arr...))

	d, err := divideFloat(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

	func(x string) {
		fmt.Println("hello", x)
	}("world")

	f := func(x string) {
		fmt.Println("hello2", x)
	}

	f("world2")

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func block11() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	r, ok := w.(ConsoleWriter)
	// r, ok := w.(io.Reader)
	if ok {
		fmt.Println("conversion ok", r)
	} else {
	}

	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")

	}

}

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

func block12() {
	// var msg = "Hello"
	// wg.Add(1)

	// go func(msg string) {
	// 	fmt.Println("Hello")
	// 	wg.Done()
	// }(msg)

	// msg = "Goodbye"

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()

	// Race conditions
	// > go run -race .
}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		// i := <-ch
		// fmt.Println(i)
		// i = <-ch
		// fmt.Println(i)

		for i := range ch {
			fmt.Println(i)
		}

		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		i := 42
		ch <- i
		i = 27
		ch <- i
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
