package main

import (
	"github.com/kogai/runner"
)

func main() {
	r := runner.New("./", "cat")
	r.Run()
}
