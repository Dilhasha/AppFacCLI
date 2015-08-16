/*
 * Copyright (c) 2015, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *   WSO2 Inc. licenses this file to you under the Apache License,
 *   Version 2.0 (the "License"); you may not use this file except
 *   in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing,
 *   software distributed under the License is distributed on an
 *   "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 *   KIND, either express or implied.  See the License for the
 *   specific language governing permissions and limitations
 *   under the License.
 */

 /* Package main handles the flow of CLI tool matching user arguments, session details to the required flags for eac command*/
package main

import (
	"os"
	"github.com/Dilhasha/AppFacCLI/cli/command"
	"github.com/codegangsta/cli"
	"github.com/Dilhasha/AppFacCLI/cli/password"
	"github.com/Dilhasha/AppFacCLI/cli/session"
	"fmt"
	"io/ioutil"
	"io"
	"encoding/json"
)

const (
	loginCommand = "login"
	//<TODO> move this to cache
	filename = "session.txt"
)
//main handles the flow of the CLI tool.
func main() {
	//Create basic cli tool
	app := cli.NewApp()
	app.Name = "appfac"
	app.Usage = "CLI Tool for WSO2 Appfactory"

	cmdFactory := command.NewFactory()
	var continueFlag bool
	var sessionObject session.Session
	var args[] string

	//command `appfac` without argument or help (h) flag
	if len(os.Args) == 1 || os.Args[1] == "help" || os.Args[1] == "h" {
		println("Showing help commands")
		app.Run(os.Args)
	}else if command, ok := cmdFactory.CheckIfCommandExists(os.Args[1]); ok {

		args = os.Args[1:]

		if(command != loginCommand) {
			//If session file exists
			if _, err := os.Stat(filename); err == nil {
				//Read session data
				data, err := ioutil.ReadFile(filename)
				if err != nil {
					panic(err)
				}
				//Get session data into a session object
				err = json.Unmarshal(data, &sessionObject)
				if(err != nil){
					println("Error occured while getting stored session.")
				}
				continueFlag = runCommand(command,args,sessionObject,cmdFactory)
			}else{
				fmt.Println("You must be logged into continue.")
				continueFlag = runCommand(loginCommand,args,sessionObject,cmdFactory)
			}
		}else {
			sessionObject = session.NewSession()
			continueFlag = runCommand(command,args,sessionObject,cmdFactory)
		}
		for(!continueFlag){
			//fmt.Println("You must be logged into continue.")
			continueFlag = runCommand(loginCommand,args,sessionObject,cmdFactory)
		}
	}else{
		println("The command you entered does not exist!")
	}

}

//getRequirements returns the list of requirements needed to run a command.
func getRequirements(command command.Command,cmdFlags []string,sessionObject session.Session,args []string)([]string){
	var i=0
	//If flag parsing is disabled
	if(command.Metadata().SkipFlagParsing){
		flags:=command.Metadata().Flags
		// Prompt for values for each flag
		requirements:= make([]string,len(flags),len(flags))
		for n := 0; n < len(flags); n++ {
			if flag, ok := flags[n].(cli.StringFlag); ok {
				if (flag.Usage != "password") {
					print(flag.Usage + " > ")
					fmt.Scanf("%s", &requirements[i])
					if(flag.Usage == "userName"){
						sessionObject.UserName=requirements[i]
					}
					i++
				}else {
					requirements[i] = password.AskForPassword("Password")
					i++
				}
			}
		}
		return requirements
	}else{
		isMatch,flagValues:= matchArgAndFlags(cmdFlags,args[1:],sessionObject)
		if(isMatch){
			return flagValues
		}else{
			return nil
		}
	}

}

//matchArgAndFlags matches the flags against user arguments and data available in session.
func matchArgAndFlags(flags []string, args []string,sessionObject session.Session) (bool,[]string) {
	var i = 0
	var requirements=make([]string,len(flags),len(flags))
	Loop:
	for _, flag := range flags {
		//Checks if flag value is present in user arguments
		containsFlag,index:=checkIfArgsContainsFlag(flag,args)
		//Checks if flag value is present in session object
		inSession,val:=checkIfInSession(flag,sessionObject)
		if (containsFlag){
			requirements[i]=args[index]
			i++
			continue Loop
		}else if(inSession){
			requirements[i]=val
			i++
			continue Loop
		}else{
			return false,requirements
		}
	}
	return true,requirements
}

//checkIfArgsContainsFlag checks whether a given flag is available in user arguments, if so returns index of flag value.
func checkIfArgsContainsFlag(flag string, args []string) (bool,int){
	for n := 0; n < len(args); n++ {
		if(args[n]==flag){
			//flag at n, so flag value at n+1
			return true,n+1
		}
	}
	return false,-1
}

//checkIfInSession returns whether the user already logged in and returns the cookie for session.
func checkIfInSession(flag string,sessionObject session.Session)(bool,string){
	if(flag=="-u"){
		return true, sessionObject.UserName
	}else if(flag=="-c"){
		return true, sessionObject.Cookie
	}
	return false, ""
}

func runCommand(commandName string , args []string, sessionObject session.Session, cmdFactory command.ConcreteFactory)(bool){
	command := cmdFactory.CmdsByName[commandName]
	cmdFlags := cmdFactory.GetCommandFlags(command)
	flagValues := getRequirements(command, cmdFlags, sessionObject, args)
	configs := cmdFactory.GetCommandConfigs(command, flagValues)
	continueFlag, cookie := command.Run(configs)
	if(commandName==loginCommand && continueFlag){
		//set session object username
		sessionObject = setSessionUserName(cmdFlags, flagValues)
		sessionObject.Cookie=cookie
		success := writeSession(sessionObject, filename)
		if(success){
			fmt.Println("Your session details are stored.")
		}else{
			fmt.Println("Error occured while storing session!")
		}
	}
	return continueFlag
}
//setSession returns a session object with userName set.
func setSessionUserName(flags []string, flagValues []string)(session.Session){
	for n := 0; n < len(flags); n++ {
		//If userName is available in flagValues set it in session
		if(flags[n]=="-u"){
			return session.Session{flagValues[n], ""}
		}
	}
	return session.Session{"",""}
}

//writeSession writes the session object to a file and returns whether it is successful.
func writeSession(sessionObject session.Session,filename string)bool{
	var n int
	var s []byte
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	s, err = json.Marshal(sessionObject)
	if err != nil {
		fmt.Println(err)
		return false
	}
	n, err = io.WriteString(file, string(s))
	if err != nil{
		fmt.Println(n, err)
		return false
	}
	file.Close()
	return true
}






