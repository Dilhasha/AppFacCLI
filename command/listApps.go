package command

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/Dilhasha/AppFacCLI/cli/formats"
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

func(applist AppList) Run(c CommandConfigs){
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
			}
		}else{
			var apps []formats.AppFormat
			err := json.Unmarshal([]byte(bodyStr), &apps)
			if(err ==nil){
				fmt.Println("You have ", len(apps)," applications")

			}

		}
	}
}
