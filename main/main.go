package main

import (
	"os"
	"github.com/Dilhasha/AppFacCLI/cli/command"
	"github.com/codegangsta/cli"
	"github.com/Dilhasha/AppFacCLI/cli/password"
	"fmt"

)

type Requirement struct {
	name string
}

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
		cmdFlags:=cmdFactory.GetCommandFlags(c)
		flagVals:=getRequirements(c,cmdFlags)
		configs:=cmdFactory.GetCommandConfigs(c,flagVals)
		c.Run(configs)
		
		

	}else{
		println("Command does not exist")
	}

}
func getRequirements(c command.Command,cmdFlags []string)([]string){

	
	var i=0
	if(c.Metadata().SkipFlagParsing){
		flags:=c.Metadata().Flags
		reqs:= make([]string,len(flags),len(flags))
		for n := 0; n < len(flags); n++ {
			if flag, ok := flags[n].(cli.StringFlag); ok {
				if (flag.Usage != "password") {
					println(flag.Usage + " > ")
					fmt.Scanf("%s", &reqs[i])
					i++
				}else {
					reqs[i] = password.AskForPassword("Password")
					i++
				}
			}
		}
		return reqs
	}else{
		isMatch,flagVals:=matchArgAndFlags(cmdFlags,os.Args[2:])
		
		if(isMatch){
			return flagVals
		}else{
			return nil
		}
	}
	
}
func matchArgAndFlags(flags []string, args []string) (bool,[]string) {
	var i=0
	var reqs=make([]string,len(flags),len(flags))

	Loop:
	for _, flag := range flags {
		containsFlag,index:=checkIfArgsContainsFlag(flag,args)
		if (containsFlag){
			reqs[i]=args[index]
			i++
			continue Loop
		}else{
			return false,reqs
		}
	}
	return true,reqs
}

func checkIfArgsContainsFlag(flag string, args []string) (bool,int){
	for n := 0; n < len(args); n++ {
		if(args[n]==flag){

			return true,n+1
		}
	}
	return false,-1
}






