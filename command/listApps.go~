package command

import (
	"fmt"
	"bytes"
)

type AppList struct {
	//to be added
}

func NewAppList() (cmd AppList) {
	return
}

func (applist AppList)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"AppList",
		Description : "Lists applications of a user",
		ShortName : "la",
		Usage:"list apps",
	}

}

func (applist AppList)Configs(reqs CommandRequirements)(configs CommandConfigs){

	var buffer bytes.Buffer
	buffer.WriteString("action=getApplicationsOfUser")

	if(reqs.UserName!=""){
		buffer.WriteString("&userName=")
		buffer.WriteString(reqs.UserName)
	}
	return CommandConfigs{
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/application/get/ajax/list.jag",
		Query:buffer.String(),
		Cookie:reqs.Cookie,
	}
}

func (applist AppList) Requirements()(reqs CommandRequirements){
	var username,cookie string
	fmt.Println("Cookie:")
	fmt.Scanf("%s", &cookie)
	fmt.Println("UserName:")
	fmt.Scanf("%s", &username)
	reqs.Cookie=cookie
	reqs.UserName=username
	return
}
