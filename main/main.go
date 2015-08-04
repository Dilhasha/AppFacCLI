package main

import (
	"os"
	"github.com/Dilhasha/AppFacCLI/cli/command"
	"github.com/codegangsta/cli"
	"github.com/Dilhasha/AppFacCLI/cli/password"
	"github.com/Dilhasha/AppFacCLI/cli/session"
	"fmt"
	"strings"
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
	var continueFlag bool=true
	var sessionObj session.Session =session.NewSession()
	var str string//store arguments temporarily

	//command `appfac` without argument
	if len(os.Args) == 1 || os.Args[1] == "help" || os.Args[1] == "h" {
		println("Showing help commands")
		app.Run(os.Args)
	}else if _, ok := cmdFactory.CmdsByName[os.Args[1]]; ok {
		for (continueFlag) {
			if(sessionObj.UserName==""){
				if(os.Args[1]!="login"){
					println("You must be logged in to continue.")
				}
				c := cmdFactory.CmdsByName["login"]
				cmdFlags := cmdFactory.GetCommandFlags(c)
				flagVals := getRequirements(c, cmdFlags,sessionObj)
				//set session obj username
				sessionObj=setSession(cmdFlags,flagVals)

				configs := cmdFactory.GetCommandConfigs(c, flagVals)
				continueFlag,sessionObj.Cookie = c.Run(configs)
			}else{
				c := cmdFactory.CmdsByName[os.Args[1]]
				cmdFlags := cmdFactory.GetCommandFlags(c)
				flagVals := getRequirements(c, cmdFlags,sessionObj)
				configs := cmdFactory.GetCommandConfigs(c, flagVals)
				continueFlag,_ = c.Run(configs)
			}
			print("appfac > ")
			fmt.Scanf("%s", &str)
			os.Args=strings.Split(str," ")
		}
	}else{
		println("Command does not exist")
	}

}
func getRequirements(c command.Command,cmdFlags []string,sessionObj session.Session)([]string){
	var i=0
	if(c.Metadata().SkipFlagParsing){
		flags:=c.Metadata().Flags
		reqs:= make([]string,len(flags),len(flags))
		for n := 0; n < len(flags); n++ {
			if flag, ok := flags[n].(cli.StringFlag); ok {
				if (flag.Usage != "password") {
					print(flag.Usage + " > ")
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
		isMatch,flagVals:=matchArgAndFlags(cmdFlags,os.Args[2:],sessionObj)
		
		if(isMatch){
			return flagVals
		}else{
			return nil
		}
	}
	
}

func matchArgAndFlags(flags []string, args []string,sessionObj session.Session) (bool,[]string) {
	var i=0
	var reqs=make([]string,len(flags),len(flags))

	Loop:
	for _, flag := range flags {
		containsFlag,index:=checkIfArgsContainsFlag(flag,args)
		inSession,val:=checkIfInSession(flag,sessionObj)
		if (containsFlag){
			reqs[i]=args[index]
			i++
			continue Loop
		}else if(inSession){
			reqs[i]=val
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

func checkIfInSession(flag string,sessionObj session.Session)(bool,string){
	if(flag=="userName"){
		return true, sessionObj.UserName
	}else if(flag=="cookie"){
		return true, sessionObj.Cookie
	}
	return false, ""
}

func setSession(flags []string, flagVals []string)(session.Session){
	for n := 0; n < len(flags); n++ {
		if(flags[n]=="userName"){
			return session.Session{flagVals[n], ""}
		}

	}
	return session.Session{"",""}
}







