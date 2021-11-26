package main

import (
	"fmt"
	"time"
)

func main() {
	type Person struct {
		Name string
		DOB  time.Time
	}
	var i *int
	var m *[2]int
	var n *string
	var p *Person

	i = new(int)
	m = new([2]int)
	n = new(string)
	p = new(Person)

	fmt.Println(i, "    ", *i)
	fmt.Println(m, "    ", *m)
	fmt.Println(n, "    ", *n)
	fmt.Println(p, "    ", *p)

	var mp *map[string]string
	mp = new(map[string]string)
	*mp = make(map[string]string) // if this omited, code will panic
	(*mp)["name"] = "lc"
	fmt.Println((*mp)["name"])
}
