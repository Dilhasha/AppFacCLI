package main

import (
	"os"
	"github.com/Dilhasha/AppFacCLI/cli/command"
	"github.com/codegangsta/cli"
	"github.com/Dilhasha/AppFacCLI/cli/password"
	"github.com/Dilhasha/AppFacCLI/cli/session"
	"fmt"
	"strings"
	"bufio"
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
	var flagVals []string
	var args[] string

	//command `appfac` without argument
	if len(os.Args) == 1 || os.Args[1] == "help" || os.Args[1] == "h" {
		println("Showing help commands")
		app.Run(os.Args)
	}else if _, ok := cmdFactory.CmdsByName[os.Args[1]]; ok {
		args=os.Args[1:]
		reader := bufio.NewReader(os.Stdin)
		for (continueFlag) {
			println("on top",args[0])
			if(sessionObj.UserName==""){
				if(args[0]!="login"){
					println("You must be logged in to continue.")
				}
				c := cmdFactory.CmdsByName["login"]
				cmdFlags := cmdFactory.GetCommandFlags(c)
				flagVals= getRequirements(c, cmdFlags,sessionObj,args)
				//set session obj username
				sessionObj=setSession(cmdFlags,flagVals)

				configs := cmdFactory.GetCommandConfigs(c, flagVals)
				continueFlag,sessionObj.Cookie = c.Run(configs)
			}else{
				println("command:"+args[0])
				c := cmdFactory.CmdsByName[args[0]]
				cmdFlags := cmdFactory.GetCommandFlags(c)
				flagVals= getRequirements(c, cmdFlags,sessionObj,args)
				configs := cmdFactory.GetCommandConfigs(c, flagVals)
				continueFlag,sessionObj.Cookie = c.Run(configs)
			}
			print("appfac > ")
			str, _ := reader.ReadString('\n')
			args=strings.Fields(str)
			println("string has :",len(args))
			for i:=0;i<len(args);i++{
				println(args[i])
			}
		}
	}else{
		println("Command does not exist")
	}

}
func getRequirements(c command.Command,cmdFlags []string,sessionObj session.Session,args []string)([]string){
	var i=0
	if(c.Metadata().SkipFlagParsing){
		flags:=c.Metadata().Flags
		reqs:= make([]string,len(flags),len(flags))
		for n := 0; n < len(flags); n++ {
			if flag, ok := flags[n].(cli.StringFlag); ok {
				if (flag.Usage != "password") {
					print(flag.Usage + " > ")
					fmt.Scanf("%s", &reqs[i])
					sessionObj.UserName=reqs[i]
					i++
				}else {
					reqs[i] = password.AskForPassword("Password")
					i++
				}
			}
		}
		return reqs
	}else{
		isMatch,flagVals:=matchArgAndFlags(cmdFlags,args[1:],sessionObj)
		
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
	if(flag=="-u"){
		return true, sessionObj.UserName
	}else if(flag=="-c"){
		return true, sessionObj.Cookie
	}
	return false, ""
}

func setSession(flags []string, flagVals []string)(session.Session){
	for n := 0; n < len(flags); n++ {
		if(flags[n]=="-u"){
			println(flags[n])
			println(flagVals[n])
			return session.Session{flagVals[n], ""}
		}

	}
	return session.Session{"",""}
}







