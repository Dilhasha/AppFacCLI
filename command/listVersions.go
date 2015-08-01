package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/Dilhasha/AppFacCLI/cli/formats"
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
		Name:"getAppVersionsInStage",
		Description : "Lists versions of an application in a stage",
		ShortName : "lv",
		Usage:"list versions",
		SkipFlagParsing:false,
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/application/get/ajax/list.jag",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "-u", Usage: "userName"},
			cli.StringFlag{Name: "-s", Usage: "stageName"},
			cli.StringFlag{Name: "-a", Usage: "applicationKey"},
			cli.StringFlag{Name: "-c", Usage: "cookie"},
		},
	}
}


func(versionsList VersionsList) Run(c CommandConfigs){
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
			var appVersions []formats.AppVersionFormat
			err := json.Unmarshal([]byte(bodyStr), &appVersions)
			if(err ==nil){
				fmt.Println("Application has ", len(appVersions[0].Versions)," versions. Details of versions are as follows.\n")
				for _, appVersion := range appVersions {
					versions:=appVersion.Versions
					for _, version := range versions{
						fmt.Println("Version:\t"+version.Version)
						fmt.Println("------------------------------------------")
						fmt.Println("autoDeploy:\t"+version.AutoDeployment)
						fmt.Println("stage:\t"+version.Stage)
						fmt.Println("isAutoBuild:\t"+version.IsAutoBuild)
						fmt.Println("isAutoDeploy:\t"+version.IsAutoDeploy)
						fmt.Println("repoURL:\t"+version.RepoURL+"\n")
					}

				}
			}

		}
	}
}
