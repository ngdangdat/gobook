package main

import (
	"flag"
	"fmt"
	"gobook/tempconv"
)

var temp = tempconv.CelciusFlag("temp", 20.0, "Input the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
