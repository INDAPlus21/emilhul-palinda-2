package main

import "fmt"

func main() {
	ch := make(chan string)
	go helperFunc(ch)
	fmt.Println(<-ch)
}

func helperFunc(ch chan<- string) {
	ch <- "Hello world!"
}
