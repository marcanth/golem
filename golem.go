package main

import (
	"./copist"
	"./zipper"
	"./utils"
	"os"
)

func main() {
	args := utils.ParseArgs(os.Args)
	zipper.Zip(args["dest"], args["zip"])
	copist.Copy(args["source"], args["dest"])
}
