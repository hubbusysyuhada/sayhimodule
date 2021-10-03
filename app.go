package gomoduleshello

import "fmt"

func SayHi() {
	defer end()
	fmt.Println("hello guys, this is method from say hi modules")
}

func init() {
	fmt.Println("initialization")
}

func end() {
	fmt.Println("stop")
	fmt.Println("thank you")
}
