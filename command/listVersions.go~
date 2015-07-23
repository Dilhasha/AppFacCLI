package command

import (
	"fmt"
	"bytes"
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

func (versionsList VersionsList) Requirements()(reqs CommandRequirements){
	var username,cookie,stageName,appKey string
	fmt.Println("Cookie:")
	fmt.Scanf("%s", &cookie)
	fmt.Println("User Name:")
	fmt.Scanf("%s", &username)
	fmt.Println("Stage Name:")
	fmt.Scanf("%s", &stageName)
	fmt.Println("Application Key:")
	fmt.Scanf("%s", &appKey)
	reqs.Cookie=cookie
	reqs.UserName=username
	reqs.Stage=stageName
	reqs.ApplicationKey=appKey
	return
}
