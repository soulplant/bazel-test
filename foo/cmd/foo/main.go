package main

import (
	"fmt"

	pb "go.sajari.com/bazeltest/foo/cmd/foo/v1"
)

func main() {
	fmt.Println("vim-go")
	fmt.Printf("%v\n", (&pb.Hello{
		Name: "hi",
	}).GetName())
}
