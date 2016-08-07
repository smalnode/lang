package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "~/Pictures/BF3/BF300.jpg"
	if info, err := os.Stat(fileName); os.IsExist(err) {
		fmt.Printf("file %s existed\n, size = %d", fileName, info.Size())
	} else {
		fmt.Printf("file %s doesn't exist\n", fileName)
	}
}
