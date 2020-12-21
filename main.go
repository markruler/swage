package main

import (
	"os"

	"github.com/markruler/swage/pkg/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
