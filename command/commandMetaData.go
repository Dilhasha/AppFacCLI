package command

import "github.com/codegangsta/cli"

type CommandMetadata struct {
	Name            string
	ShortName       string
	Usage           string
	Url				string
	Description     string
	Flags           []cli.Flag
	SkipFlagParsing bool
}

