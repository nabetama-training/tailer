package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

type TargetFile struct {
	FilePath string
	Follow   bool
}

func (tf *TargetFile) Tail() {
	t, err := tail.TailFile(tf.FilePath, tail.Config{Follow: tf.Follow})
	PanicIf(err)

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
