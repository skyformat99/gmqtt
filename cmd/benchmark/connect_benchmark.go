package main

import (
	"fmt"
	"github.com/DrmagicE/gmqtt/cmd/benchmark/connect"
	"os"
)

func main() {
	cmd := &connect.Command{}
	if err := cmd.Run(os.Args[1:]...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
