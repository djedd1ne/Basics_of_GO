package main

//global variable
var url = "https://djmekki.com"

var price = 3.1

func main() {
	//function-scoped variables
	message := "Hello from Go"
	var isReady bool

	print(message, price, url, isReady)
}