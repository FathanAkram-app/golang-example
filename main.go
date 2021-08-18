package main

import (
	"errors"
	"fmt"
	"math"
)

// stack memory
// doesnt need garbage collector, self cleaning
func main() {

	// ":="  = initial variable / allocate a new memory
	i, j := 42, 100

	// variable initial value = 0
	var x int
	fmt.Println(x)

	// array
	// fixed index, not appendable
	array := [5]int{0, 1, 2, 3, 4}
	array[2] = 7
	fmt.Println(array)

	// a sliced array
	// its appendable
	slicedArray := []int{0, 1, 7}
	slicedArray = append(slicedArray, 2)
	fmt.Println(slicedArray)

	// map
	vertices := make(map[string]int)
	vertices["apple"] = 3
	vertices["grape"] = 5
	vertices["watermelon"] = 1
	fmt.Println(vertices)
	fmt.Printf("there are %v grapes \n", vertices["grape"])
	delete(vertices, "grape")

	// "=" = set variable
	j = 200

	// & = memory address of variable
	fmt.Println(&i, &j, j)

	// "*" = value of memory address / dereference
	p := &i
	fmt.Println(*p, p)
	*p = 21
	fmt.Println(*p, p)

	// for loop
	for i := 0; i < len(slicedArray); i++ {
		// if and else statement
		if slicedArray[i] <= 0 {
			fmt.Println("is less than zero")
		} else if slicedArray[i]%2 == 0 {
			fmt.Printf("%v is even \n", slicedArray[i])
			slicedArray = append(slicedArray, i)

		} else {
			fmt.Printf("%v is odd \n", slicedArray[i])
		}
	}
	// foreach
	for index, value := range array {
		fmt.Println("index:", index, "value:", value)

	}

	// while loop
	count := 1
	for count <= 3 {
		fmt.Printf("counted %v from while loop\n", count)
		count++
	}

	p = &j
	*p = *p / 2
	fmt.Println(j, &j)
	check()

	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	concurrencyTest()
}

// math square root, error handling
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined negative numbers")
	}
	return math.Sqrt(x), nil
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
