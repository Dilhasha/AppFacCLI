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
	"strings"
	"io/ioutil"
	"net/http"
	"github.com/codegangsta/cli"

)

/* Login is the implementation of the command to log into a user account in app factory */
type Login struct {
	Url string
}

func NewLogin(url string) (cmd Login) {
	return Login{
		Url:url,
	}
}

/* Returns metadata for login.*/
func (login Login)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"login",
		Description : "Login to app factory",
		ShortName : "l",
		Usage:"login",
		Url:login.Url,
		SkipFlagParsing:true,
		Flags: []cli.Flag{
			cli.StringFlag{Name:"-u",Usage:"userName"},
			cli.StringFlag{Name: "-p", Usage: "password"},
		},
	}
}

/* Run calls the Run function of CommandConfigs and verifies the response from that call.*/
func(login Login) Run(c CommandConfigs)(bool,string){
	var resp *http.Response
	var bodyStr string
	resp=c.Run()
	defer resp.Body.Close()

	if(resp.Status=="200 OK"){
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr=string(body)
		println(bodyStr)
		if(strings.Contains(bodyStr, "true")){
			fmt.Println("You have Successfully logged in.")
			cookie:=strings.Split(resp.Header.Get("Set-Cookie"),";")
			c.Cookie=cookie[0]
		}else{
			fmt.Println("Authorization failed. Please try again!")
			return false,c.Cookie
		}
	}
	return true,c.Cookie
}

