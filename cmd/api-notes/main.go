package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/reddec/api-notes/cmd/api-notes/commands"
)

//nolint:gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

type Config struct {
	Serve commands.ServeCMD `command:"serve" alias:"run" description:"Run server"`
}

func main() {
	var config Config
	parser := flags.NewParser(&config, flags.Default)
	parser.ShortDescription = "API-Notes"
	parser.LongDescription = fmt.Sprintf("Serve notes with API\napit-notes %s, commit %s, built at %s by %s\nAuthor: Aleksandr Baryshnikov <owner@reddec.net>", version, commit, date, builtBy)
	parser.EnvNamespace = "API_NOTES"
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
}
