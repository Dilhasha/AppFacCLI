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
package command

import (
	"bytes"
	"errors"
	"github.com/codegangsta/cli"
	"github.com/Dilhasha/AppFacCLI/cli/urls"
	"io/ioutil"
	"os"
	"encoding/json"
)

const (
	//Keyword for starting a query
	queryStarter = "action"
	connector = "&"
	equator = "="
)

type ConcreteFactory struct {
	CmdsByName map[string]Command
}


/*GetByCmdName returns command given the command name or short name*/
func (factory ConcreteFactory) GetByCmdName(cmdName string) (cmd Command, err error) {
	cmd, found := factory.CmdsByName[cmdName]
	if !found {
		for _, command := range factory.CmdsByName {
			if command.Metadata().ShortName == cmdName {
				return command, nil
			}
		}
		err = errors.New("Command not found")
	}
	return
}

/* NewFactory returns a new concreteFactory with with a map of commands.*/
func NewFactory() (factory ConcreteFactory) {
	//Get Urls
	urls := urls.GetUrls()
	//Create map of commands
	factory.CmdsByName = make(map[string]Command)
	factory.CmdsByName["login"] = NewLogin(urls.Login)
	factory.CmdsByName["listApps"] = NewAppList(urls.ListApps)
	factory.CmdsByName["listVersions"] = NewVersionsList(urls.ListVersions)
	factory.CmdsByName["createApp"] = NewAppCreation(urls.CreateApp)
	factory.CmdsByName["exit"] = NewExit(urls.Exit)
	factory.CmdsByName["getAppInfo"] = NewAppInfo(urls.GetAppInfo)
	factory.CmdsByName["buildApp"] = NewArtifact(urls.CreateArtifact)
	factory.CmdsByName["getBuildSuccessInfo"] = NewBuildSuccessInfo(urls.GetBuildSuccessInfo)
	factory.CmdsByName["printLogs"] = NewPrintLogs(urls.PrintLogs)
	return
}

/* GetCommandFlags converts flags into a list of strings.*/
func (factory ConcreteFactory) GetCommandFlags(command Command) []string {
	var flags []string
	for _, flag := range command.Metadata().Flags {
		switch flagType := flag.(type) {
		default:
		case cli.StringFlag:
			flags = append(flags, flagType.Name)
		}
	}
	return flags
}

/* GetCommandConfigs returns a CommandConfigs struct based on flags nd flag values.*/
func (factory ConcreteFactory) GetCommandConfigs(command Command,flagValues []string) CommandConfigs {

	var buffer bytes.Buffer
	flags:=command.Metadata().Flags
	var cookie string

	buffer.WriteString(queryStarter + equator + command.Metadata().Name)

	for n := 0; n < len(flags); n++ {
		//If flag is a string flag
		if flag, ok := flags[n].(cli.StringFlag); ok {
			if(flag.Usage == "cookie"){
				cookie = flagValues[n]
			}else{
				buffer.WriteString(connector + flag.Usage + equator)
				buffer.WriteString(flagValues[n])
			}
		}
	}
	query := buffer.String()

	return CommandConfigs{
		Url:command.Metadata().Url,
		Query:query,
		Cookie:cookie,
	}
}

/*Get url values from file to a Urls object*/
func getURLs(filename string)(urls.Urls){
	var urlValues urls.Urls
	if _ , err := os.Stat(filename); err == nil {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic (err)
		}
		err = json.Unmarshal(data , &urlValues)
		if (err != nil) {
			panic (err)
		}
	}
	return urlValues
}
