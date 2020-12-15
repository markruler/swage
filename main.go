package main

import (
	"fmt"

	"github.com/markruler/swage/pkg/cmd"
)

var tmp string

func init() {
	tmp = "swage"
}

func main() {
	fmt.Printf("%s\n", cmd.Print(tmp))
}
