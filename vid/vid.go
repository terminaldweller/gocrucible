package main

import (
	"fmt"
	"sync"
)

// import "time"

var k int = 23

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter int = 0

func main() {
	// var i int = 42
	// j := 33
	// i := 34.
	// i = float64(j)
	// var x complex128 = 0
	// // var i1 int8
	// fmt.Println(i)
	// fmt.Printf("%v, %T", k, k)
	// fmt.Printf("%v, %T", x, x)
	// const (
	// 	blue  = iota
	// 	red   = iota
	// 	green = iota
	// )

	// const (
	// 	a = iota
	// 	b = iota
	// 	c = iota
	// )

	// var name [3]string
	// fmt.Printf("%v, %T", name, name)

	// var magic [8]byte
	// magic[0] = 23
	// fmt.Printf("%v, %T", magic, magic)

	// fmt.Printf("%v, %T\n", blue, blue)
	// fmt.Printf("%v, %T\n", red, red)
	// fmt.Printf("%v, %T\n", green, green)
	// fmt.Printf("%v, %T\n", a, a)
	// fmt.Printf("%v, %T\n", b, b)
	// fmt.Printf("%v, %T\n", c, c)

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// // b := a[:]
	// // c := a[3:]
	// // d := a[:6]
	// // e := a[3:6]
	// f := a[3:]
	// fmt.Println(f)

	// a := make([]int, 3, 100)
	// a = append(a, 1)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// a = append(a[:2], a[3:]...)

	// fmt.Println(a)

	// type Color struct {
	// 	R          byte
	// 	G          byte
	// 	B          byte
	// 	ColorNames []string
	// }

	// NavyBlue := Color{
	// 	R:          0xff,
	// 	G:          0xff,
	// 	B:          0xff,
	// 	ColorNames: []string{"blud", "darkblue"}}

	// fmt.Println(NavyBlue)

	// type Animal struct {
	// 	Name   string
	// 	Origin string
	// }

	// type Bird struct {
	// 	Animal
	// 	Name     string
	// 	SpeedKPH float32
	// 	CanFly   bool
	// }

	// b := Bird{}

	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// b.Name = "Emu"
	// b.Animal.Name = "fuck"
	// fmt.Println(b.Name)
	// fmt.Println("loco")
	// defer fmt.Println("lupo")
	// defer fmt.Println("vida")

	// a := [6]int{1, 2, 3, 4, 5, 6}
	// b := &a[0]
	// c := &a[1]
	// d := &a
	// fmt.Println(a, b, c, d)
	// var ms *myStruct
	// ms = new(myStruct)
	// ms.A = 123
	// fmt.Println(ms)
	// for i := 0; i < 5; i++ {
	// 	sayMessage("hellog go!", i)
	// }
	// sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// a := sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// println(*a)

	// var divvy func(float64, float64) (float64, error)
	// divvy = func(a, b float64) (float64, error) {
	// 	if b == 0.0 {
	// 		return 0.0, fmt.Errorf("cannot divide by zero")
	// 	} else {
	// 		return a / b, nil
	// 	}
	// }

	// fmt.Println(divvy(2.3, 3.4))

	// func() {
	// 	msg := "yo yo yo"
	// 	fmt.Println(msg)
	// }()

	// g := greeter{
	// 	greeting: "hello",
	// 	name:     "go",
	// }

	// g.greet()
	// fmt.Println("the new name is", g.name)
	// var w Writer = ConsoleWriter{}
	// w.Write([]byte("Hello Go!"))
	// go sayHello()

	// var msg = "hello"
	// wg.Add(1)
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	// msg = "goodbye"
	// wg.Wait()
	// time.Sleep(100 * time.Millisecond)
	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	m.RLock()
	// 	go sayHello()
	// 	m.Lock()
	// 	go increment()
	// }
	// wg.Wait()
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}

// type myStruct struct {
// 	A    int
// 	B    int
// 	C    int
// 	Name []string
// }

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("the value of the index is", idx)
}

func sayGreeting(greeting, name string) {

}

func sum(values ...int) {
	result := 0
	for _, v := range values {
		result += v
	}

	fmt.Println(result)
}

func sum2(values ...int) *int {
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func fdiv(a, b float64, c int) float64 {
	return a + b
}

// func sayHello() {
// 	fmt.Println("Hello")
// }

func sayHello() {
	fmt.Printf("hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
