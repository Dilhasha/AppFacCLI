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
	"io/ioutil"
	"net/http"
	"github.com/codegangsta/cli"
)

/* PrintLogs is the implementation of the command to display build logs of a given application of app factory user */
type PrintLogs struct {
	Url string
}

func NewPrintLogs(url string) (cmd PrintLogs) {
	return PrintLogs{
		Url:url,
	}
}

/* Returns metadata for printing build logs.*/
func (printLogs PrintLogs)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"printBuildLogs",
		Description : "print logs for a given build",
		ShortName : "pl",
		Usage:"print logs",
		Url:printLogs.Url,
		SkipFlagParsing:false,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "-a", Usage: "applicationKey"},
			cli.StringFlag{Name: "-v", Usage: "applicationVersion"},
			cli.StringFlag{Name: "-d", Usage: "tenantDomain"},
			cli.StringFlag{Name: "-i", Usage: "lastBuildId"},
			cli.StringFlag{Name: "-r", Usage: "forkedRepository"},
			cli.StringFlag{Name: "-c", Usage: "cookie"},
		},
	}
}

/* Run calls the Run function of CommandConfigs and verifies the response from that call.*/
func(printLogs PrintLogs) Run(c CommandConfigs)(bool,string){
	var resp *http.Response
	var bodyStr string
	resp = c.Run()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if (resp.Status == "200 OK") {
		bodyStr = string(body)
		println(bodyStr)


	}
	return true,c.Cookie
}
