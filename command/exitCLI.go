package command

import (
	"fmt"
	"github.com/codegangsta/cli"
)

type exitCLI struct {
	//to be added
}

func NewExit() (cmd exitCLI) {
	return
}

func (exit exitCLI)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"exit",
		Description : "Exists the CLI",
		ShortName : "ex",
		Usage:"exit tool",
		Url:"",
		SkipFlagParsing:true,
		Flags: []cli.Flag{},
	}

}

func(exit exitCLI) Run(c CommandConfigs)(bool,string){
	fmt.Println("Exiting Appfac CLI..")
	return false,c.Cookie
}
