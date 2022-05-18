package main

import "fmt"

func Run() error {
	fmt.Println("Starting up API")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
