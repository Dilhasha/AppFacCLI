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
	"github.com/codegangsta/cli"
)

/* createArtifact is the implementation of the command to create an artifact and deploy it in app factory */

type Artifact struct {
	Url string
}

func NewArtifact(url string) (cmd Artifact) {
	return Artifact{
		Url:url,
	}
}

/* Returns metadata for artifact creation*/
func (artifact Artifact)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"createArtifact",
		Description : "create artifact of an application",
		ShortName : "ca",
		Usage:"create artifact",
		Url:artifact.Url,
		SkipFlagParsing:false,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "-a", Usage: "applicationKey"},
			cli.StringFlag{Name: "-c", Usage: "cookie"},
			cli.StringFlag{Name: "-v", Usage: "version"},
			cli.StringFlag{Name: "-rv", Usage: "revision"},
			cli.StringFlag{Name: "-d", Usage: "doDeploy"},
			cli.StringFlag{Name: "-s", Usage: "stage"},
			cli.StringFlag{Name: "-t", Usage: "tagName"},
			cli.StringFlag{Name: "-r", Usage: "repoFrom"},
		},
	}
}

/* Run calls the Run function of CommandConfigs and verifies the response from that call.*/
func(artifact Artifact) Run(configs CommandConfigs)(bool,string){
	response := configs.Run()
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if (response.Status == "200 OK") {
		bodyString := string(body)
		println(bodyString)
	}
	return true,configs.Cookie
}
