package main

import "fmt"

type Person struct {
	name string
	age  int
	pet  string
}

type Bicycle struct {
	brand  string
	typ    string // racing, mountainbike, bmx
	weight float32
}

func main() {
	people()
	bike()
}

func people() {
	fmt.Println("--------- people --------\n")
	peter := Person{
		name: "Peter Parke",
		age:  23,
		pet:  "Spider",
	}

	fmt.Println(peter.pet)
	fmt.Println(peter.name)
	fmt.Println(peter.age)
	//bike()
}

func bike() {
	fmt.Println("--------- bike --------\n")

	bike := Bicycle{
		brand:  "Canondale",
		typ:    "BMX",
		weight: 7.2,
	}

	fmt.Println(bike.brand)
	fmt.Println(bike.typ)
	fmt.Println()
}
