package test

import "fmt"

func TestHelper() {
	fmt.Println("TestHelper v1.1")
}

func SayHello(name string) {
	fmt.Printf("%s: Hello!", name)
}

func SayHelloBack(name string) string {
	return fmt.Sprintf("%s: Hello!", name)
}
