package main

import (
	"os"

	"github.com/carlmjohnson/exitcode"
	"https://github.com/izner32/deployer/cmd"
)

func main() {
	exitcode.Exit(Deployer.CLI(os.Args[1:]))
}
