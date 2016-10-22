package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"os"
)

func main() {
	t, err := tail.TailFile("/var/log/nginx/access.log", tail.Config{Follow: true})
	PanicIf(err)

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

func PanicIf(err error) {
	if err != nil {
		fmt.Errorf("Error: %s", err)
		os.Exit(1)
	}
}
