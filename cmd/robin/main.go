package main

import (
	"log"

	"github.com/loeffel-io/robin-bazel/internal/calc"
)

func main() {
	log.Printf("%d", calc.Add(1, 2))
}
