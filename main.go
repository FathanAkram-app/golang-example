package main

import "fmt"

// stack memory
// doesnt need garbage collector, self cleaning
func main() {
	// ":="  = initial variable / allocate a new memory
	i, j := 42, 100
	// "=" = set variable
	j = 200
	// & = memory address of variable
	fmt.Println(&i, &j, j)
	// "*" = value of memory address
	p := &i
	fmt.Println(*p, p)
	*p = 21
	fmt.Println(*p, p)

	p = &j
	*p = *p / 2
	fmt.Println(j, &j)
	check()
}

// structure / object
type person struct {
	name string
	age  int
}

// heap memory
// it affect performance because of the garbage collector
func initPerson() *person {
	m := person{name: "noname", age: 10}
	fmt.Printf("initPerson --> %p\n", &m)
	return &m
}

func check() {
	fmt.Printf("main --> %p\n", initPerson())
}
