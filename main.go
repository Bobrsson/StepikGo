package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)
	fmt.Print(a[2])
	fmt.Println(string(a[len(a)-1]))

}