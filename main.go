package main

import (
	"./copist"
	"./zipper"
	"flag"
	"os"
	"time"
)

func main() {
	var out string
	flag.StringVar(&out, "out", time.Now().Format("20060102150405"), "Enter with the zip name here")

	zipper.Zip(os.Args[2], out)
	copist.Copy(os.Args[1], os.Args[2])
}
