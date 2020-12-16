package main

import (
	"log"

	"github.com/markruler/swage/pkg/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		return
	}
}
