package main

import (
	"./copist"
	"./zipper"
	"./util"
	"flag"
	"os"
	"time"
)

func main() {
	args := utils(as.Args)
	zipper.Zip(args["dest"], args["zip"])
	copist.Copy(args["source"], args["dest"])
}
