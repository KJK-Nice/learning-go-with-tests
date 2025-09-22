package main

import (
	"os"
	"time"

	clockface "learning-go-with-tests/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
