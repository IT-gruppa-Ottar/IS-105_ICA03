package main

import (
	"./fileutils"
	"fmt"
)

func main() {
	fmt.Println("hex for lang01.wl")
	fmt.Printf("% X", fileutils.FileToByteslice("lang01.wl"))
	fmt.Println()

	fmt.Println("hex for lang02.wl")
	fmt.Printf("% X", fileutils.FileToByteslice("lang02.wl"))
	fmt.Println()

	fmt.Println("hex for lang03.wl")
	fmt.Printf("% X", fileutils.FileToByteslice("lang03.wl"))
}

