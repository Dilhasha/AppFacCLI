package main

import (
	"os"
	"github.com/Dilhasha/AppFacCLI/cli/command"
	"github.com/codegangsta/cli"
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
	}else if _, ok := cmdFactory.CmdsByName[os.Args[1]]; ok {
		c := cmdFactory.CmdsByName[os.Args[1]]
		//requirements:=c.Requirements()
		//configs:=c.Configs(requirements)
		configs := c.Configs(c.Metadata().Flags)
		c.Run(configs);
	}

}
func matchArgAndFlags(flags []string, args []string) string {
	var badFlag string
	var lastPassed bool
	multipleFlagErr := false

Loop:
for _, arg := range args {
	prefix := ""

	//only take flag name, ignore value after '='
	arg = strings.Split(arg, "=")[0]

	if arg == "--h" || arg == "-h" {
		continue Loop
	}

	if strings.HasPrefix(arg, "--") {
		prefix = "--"
	} else if strings.HasPrefix(arg, "-") {
		prefix = "-"
	}
	arg = strings.TrimLeft(arg, prefix)

	//skip verification for negative integers, e.g. -i -10
	if lastPassed {
		lastPassed = false
		if _, err := strconv.ParseInt(arg, 10, 32); err == nil {
			continue Loop
		}
	}

	if prefix != "" {
		for _, flag := range flags {
			if flag == arg {
				lastPassed = true
				continue Loop
			}
		}

		if badFlag == "" {
			badFlag = fmt.Sprintf("\"%s%s\"", prefix, arg)
		} else {
			multipleFlagErr = true
			badFlag = badFlag+fmt.Sprintf(", \"%s%s\"", prefix, arg)
		}
	}
}

	if multipleFlagErr && badFlag != "" {
		badFlag = fmt.Sprintf("%s %s", T("Unknown flags:"), badFlag)
	} else if badFlag != "" {
		badFlag = fmt.Sprintf("%s %s", T("Unknown flag"), badFlag)
	}

	return badFlag
}








