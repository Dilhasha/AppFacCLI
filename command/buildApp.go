package command

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/Dilhasha/AppFacCLI/cli/formats"
	"github.com/codegangsta/cli"
)

type Build struct {
	//should get repo location for login details
}

func NewBuild() (cmd Build) {
	return
}

func (build Build)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"Build",
		Description : "Trigger a build for a given app",
		ShortName : "b",
		Usage:"build",
		SkipFlagParsing:false,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u", Usage: "userName"},
			cli.StringFlag{Name: "s", Usage: "stageName"},
			cli.StringFlag{Name: "d", Usage: "deployAction"},
			cli.StringFlag{Name: "a", Usage: "applicationKey"},
			cli.StringFlag{Name: "c", Usage: "cookie"},
			cli.StringFlag{Name: "t", Usage: "tagName"},
		},
	}

}

func (build Build)Configs(reqs CommandRequirements)(configs CommandConfigs){

	var buffer bytes.Buffer
	buffer.WriteString("action=deployArtifact")
	if(reqs.ApplicationKey!=""){
		buffer.WriteString("&applicationKey=")
		buffer.WriteString(reqs.ApplicationKey)
	}
	if(reqs.Stage!=""){
		buffer.WriteString("&stage=")
		buffer.WriteString(reqs.Stage)
	}
	if(reqs.Version!=""){
		buffer.WriteString("&version=")
		buffer.WriteString(reqs.Version)
	}
	if(reqs.TagName!=""){
		buffer.WriteString("&tagName=")
		buffer.WriteString(reqs.TagName)
	}
	if(reqs.DeployAction!=""){
		buffer.WriteString("&deployAction=")
		buffer.WriteString(reqs.DeployAction)
	}
	return CommandConfigs{
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/build/add/ajax/add.jag",
		Query:buffer.String(),
		Cookie:reqs.Cookie,
	}
}

func (build Build) Requirements(args []string)(reqs CommandRequirements){
	if(!build.Metadata().SkipFlagParsing){
		reqs.Cookie=args[0]
		reqs.ApplicationKey=args[1]
		reqs.DeployAction=args[2]
		reqs.Stage=args[3]
		reqs.TagName=args[4]
		reqs.Version=args[5]
	}
	return
}
func(build Build) Run(c CommandConfigs) {
	var resp *http.Response
	var bodyStr string
	resp = c.Run()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if (resp.Status == "200 OK") {
		bodyStr = string(body)
		var errorFormat formats.ErrorFormat
		err := json.Unmarshal([]byte(bodyStr), &errorFormat)
		if (err == nil) {
			//<TODO> Make these error checking functionality common
			if (errorFormat.ErrorCode == http.StatusUnauthorized) {
				fmt.Println("Your session has expired.Please login and try again!")
			}else if (strings.Contains(bodyStr, "null")) {
				fmt.Println("Build is being triggered.....")
			}
		}
	}
}





