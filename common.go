package main

import (
	"fmt"
	"os"
)

func PanicIf(err error) {
	if err != nil {
		fmt.Errorf("Error: %s", err)
		panic(err)
		os.Exit(1)
	}
}
