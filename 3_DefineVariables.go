package main

import "fmt"

func main (){
	var x int 
	var name string //nil by default
	const y = 2 //const only bool, string, numbers

	var z int = 2 //initilize with creation

	var text string 
	text = "Hello!"

	otherText := "Bye!" // initilization shortcut

	fmt.Println(x,name,y,z,text,otherText)
}