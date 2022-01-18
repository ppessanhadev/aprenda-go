package main

import "fmt"


func main() {
	number_of_bytes, error := fmt.Println("Hello world", "Oi galera!", 100)
	fmt.Println(number_of_bytes, error)
}
