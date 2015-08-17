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
	"io/ioutil"
	"net/http"
	"github.com/codegangsta/cli"
	"github.com/Dilhasha/AppFacCLI/cli/formats"
	"encoding/json"
)

/* AppCreation is the implementation of the command to create an application in App factory. */

type AppCreation struct {
	Url string
}

func NewAppCreation(url string) (cmd AppCreation) {
	return AppCreation{
		Url : url,
	}
}

/* Returns metadata for application creation*/
func (appCreation AppCreation)Metadata() CommandMetadata{
	return CommandMetadata{
		Name : "createNewApplication",
		Description : "Creates a new application",
		ShortName : "cap",
		Usage : "create application",
		Url : appCreation.Url,
		SkipFlagParsing : false,
		Flags : []cli.Flag{
			cli.StringFlag{Name: "-u", Usage: "userName"},
			cli.StringFlag{Name: "-c", Usage: "cookie"},
			cli.StringFlag{Name: "-k", Usage: "applicationKey"},
			cli.StringFlag{Name: "-n", Usage: "applicationName"},
			cli.StringFlag{Name: "-d", Usage: "applicationDescription"},
			cli.StringFlag{Name: "-t", Usage: "appType"},
			cli.StringFlag{Name: "-r", Usage: "repositoryType"},
		},
	}
	
}

/* Run calls the Run function of CommandConfigs and verifies the response from that call.*/
func(appCreation AppCreation) Run(configs CommandConfigs) (bool,string){
	var response *http.Response
	var bodyString string
	response = configs.Run()
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if (response.Status == "200 OK") {
		bodyString = string(body)
		var errorFormat formats.ErrorFormat
		err := json.Unmarshal([]byte(bodyString), &errorFormat)

		if (err == nil) {
			//<TODO> Refine error checking functionality
			if (errorFormat.ErrorCode == http.StatusUnauthorized) {
				fmt.Println("Your session has expired.Please login and try again!")
				return false, configs.Cookie
			}
		}
		var successMessage formats.SuccessFormat
		err = json.Unmarshal([]byte(bodyString), &successMessage)
		if(err == nil){
			fmt.Println(successMessage.Message)
		}
	}
	return true , configs.Cookie
}
