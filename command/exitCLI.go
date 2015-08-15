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
	"fmt"
	"github.com/codegangsta/cli"
)

type exitCLI struct {
	Url string
}

/* exitCLI is the implementation of the command to exit app factory */

func NewExit(url string) (cmd exitCLI) {
	return exitCLI{
		Url:url,
	}
}

/* Returns metadata for exit*/
func (exit exitCLI)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"exit",
		Description : "Exists the CLI",
		ShortName : "ex",
		Usage:"exit tool",
		Url:exit.Url,
		SkipFlagParsing:true,
		Flags: []cli.Flag{},
	}

}

/* Run calls the Run function of CommandConfigs and verifies the response from that call.*/
func(exit exitCLI) Run(configs CommandConfigs)(bool,string){
	fmt.Println("Exiting Appfac CLI..")
	return true,configs.Cookie
}
