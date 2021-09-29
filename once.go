package main

import (
	"fmt"
	"sync"
)
var a string
var once sync.Once
func setup() {
	a = "hello, world"
}
func doPrint() { 
	once.Do(setup)
	fmt.Printf(a) 
}
func twoprint() { 
	go doPrint() 
	go doPrint()
}