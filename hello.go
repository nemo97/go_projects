package main

import (
	"encoding/json"
	"fmt"
)

// type SampleStruc struce {

// }
//func Marshal(v interface{}) ([]byte, error)

const m string = "Constant"

func plus(a int, b int) int {

	return a + b
}

func tupple(a int, b int) (int, int) {

	return a, b
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type geomatry interface {
	area() float64
	perim() float64
}

type rect struct {
	w, h float64
}

type cicle struct {
	d float64
}

func (r rect) area() float64 {

	return r.w * r.h
}

func (r rect) perim() float64 {
	return 2 * (r.w + r.h)
}

func (c cicle) area() float64 {

	return 2 * 3.14 * (c.d / 2) * (c.d / 2)
}
func (c cicle) perim() float64 {

	return 2 * 3.14 * (c.d / 2)
}

func main() {
	var varInt int16 = 10
	f := 10
	fmt.Println("Hello World " + m)
	fmt.Printf("Hello Print %d f= %d", varInt, f)

	/// int array
	a := make([]int, 10)
	var b [10]int

	i := 0
	for i < 10 {
		a[i] = i * i
		b[i] = plus(i, i)
		i++
		fmt.Printf("Iterate %d \n", i)
	}

	fmt.Println("s=et ", a, b)

	/// map

	k := make(map[string]int)

	k["k1"] = 2
	k["k2"] = 3

	fmt.Println("map ", k)

	for key, val := range k {
		fmt.Printf("\n Key %s val %d", key, val)
	}

	x, y := tupple(3, 4)

	fmt.Printf("\n Testing multiple return %d %d", x, y)

	// paying with Struct\

	person := Person{}

	person.Name = "Test name "
	person.Age = 30

	fmt.Printf("Priting '%s' ", person.Name)

	var g1 geomatry = rect{h: 10, w: 10}

	fmt.Printf("\n rect area %f ", g1.area())

	var g2 geomatry = cicle{d: 10}

	fmt.Printf("\n cicle area %f ", g2.area())

	p1 := Person{Name: "Subhas", Age: 34}

	j1, _ := json.Marshal(p1)
	fmt.Printf("json %s", j1)
}
