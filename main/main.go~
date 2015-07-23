package main

import (
	"os"
	"github.com/Dilhasha/AppFacCLI/cli/command"
	"github.com/codegangsta/cli"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "appfac"
	app.Usage = "CLI Tool for WSO2 Appfactory"
	app.Action = func(c *cli.Context) {
		println("first appfac CLI command!")
	}
	cmdFactory := command.NewFactory()

	//app.Run(os.Args)

	//command `appfac` without argument
	if len(os.Args) == 1 || os.Args[1] == "help" || os.Args[1] == "h" {
		println("Showing help commands")
		app.Run(os.Args)
	}else if _, ok := cmdFactory.CmdsByName[os.Args[1]]; ok{
		c:=cmdFactory.CmdsByName[os.Args[1]]
		requirements:=c.Requirements()
		fmt.Println(requirements)
		configs:=c.Configs(requirements)
		configs.Run();
	}

}








