package command

import (
	"bytes"
	"github.com/codegangsta/cli"
)

type VersionsList struct {
	//to be added
}

func NewVersionsList() (cmd VersionsList) {
	return
}

func (versionsList VersionsList)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"VersionsList",
		Description : "Lists versions of an application in a stage",
		ShortName : "lv",
		Usage:"list versions",
		SkipFlagParsing:false,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u", Usage: "userName"},
			cli.StringFlag{Name: "s", Usage: "stageName"},
			cli.StringFlag{Name: "a", Usage: "applicationKey"},
			cli.StringFlag{Name: "c", Usage: "cookie"},
		},
	}
}

func (versionsList VersionsList)Configs(reqs CommandRequirements)(configs CommandConfigs){

	var buffer bytes.Buffer
	buffer.WriteString("action=getAppVersionsInStage")

	if(reqs.UserName!=""){
		buffer.WriteString("&userName=")
		buffer.WriteString(reqs.UserName)
	}
	if(reqs.Stage!=""){
		buffer.WriteString("&stageName=")
		buffer.WriteString(reqs.Stage)
	}
	if(reqs.ApplicationKey!=""){
		buffer.WriteString("&applicationKey=")
		buffer.WriteString(reqs.ApplicationKey)
	}
	return CommandConfigs{
		Url:"https://apps.cloud.wso2.com//appmgt/site/blocks/application/get/ajax/list.jag",
		Query:buffer.String(),
		Cookie:reqs.Cookie,
	}
}

func (versionsList VersionsList) Requirements(args []string)(reqs CommandRequirements){
	if(!versionsList.Metadata().SkipFlagParsing){
		reqs.Cookie=args[0]
		reqs.UserName=args[1]
		reqs.Stage=args[2]
		reqs.ApplicationKey=args[3]
	}
	return
}
func(versionsList VersionsList) Run(c CommandConfigs){

}
