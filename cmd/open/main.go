package main

import (
	"os"

	"github.com/tanema/open"
)

func main() {
	open.URL(os.Args[1])
}
