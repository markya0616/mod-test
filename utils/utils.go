package utils

import (
	"fmt"

	_ "google.golang.org/grpc"
)

func PrintHello() {
	fmt.Println("Hello world!!")
}

func PrintHelloV2() {
	fmt.Println("Hello world v2!!")
}
