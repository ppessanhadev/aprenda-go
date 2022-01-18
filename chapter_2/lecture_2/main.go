package main

import "fmt"


func main() {
	_, error := fmt.Println("Hello world", "Oi galera!", 100)
	fmt.Println(error)
}
