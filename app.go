package gomoduleshello

import "fmt"

func SayHi() {
	defer end()
	fmt.Println("hello guys, this is method from say hi modules")
	fmt.Println("new version of 1.0.2")
}

func init() {
	fmt.Println("initialization")
}

func end() {
	fmt.Println("stop")
	fmt.Println("thank you")
}
