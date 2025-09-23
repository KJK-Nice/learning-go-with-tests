package main

import (
	"os"
	"time"

	"learning-go-with-tests/maths/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
