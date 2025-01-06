package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]bool

// init is the only exception for func overloading
func init() { 
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[8] = "Algeria"
	fmt.Println("Countries saved")
}