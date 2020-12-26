package main

import (
	"os"

	"github.com/cxsu/swage/pkg/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
