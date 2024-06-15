package main

import (
	"github.com/alecthomas/kong"
)

var (
	version = "dev"

	cli struct {
		Version kong.VersionFlag `name:"version" help:"Print version information and quit"`
	}
)

func main() {
	kong.Parse(&cli, kong.Vars{"version": version})
}
