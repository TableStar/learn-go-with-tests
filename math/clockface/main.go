package main

import (
	"os"
	"time"

	clockface "github.com/tablestar/learn-go-with-tests/math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
