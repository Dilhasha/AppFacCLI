package command

import (
	"fmt"
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
		Name:"deployArtifact",
		Description : "Trigger a build for a given app",
		ShortName : "b",
		Usage:"build ",
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/build/add/ajax/add.jag",
		SkipFlagParsing:false,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "-u", Usage: "userName"},
			cli.StringFlag{Name: "-s", Usage: "stageName"},
			cli.StringFlag{Name: "-d", Usage: "deployAction"},
			cli.StringFlag{Name: "-a", Usage: "applicationKey"},
			cli.StringFlag{Name: "-c", Usage: "cookie"},
			cli.StringFlag{Name: "-t", Usage: "tagName"},
		},
	}

}



func(build Build) Run(c CommandConfigs) (bool,string){
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
	return true,c.Cookie
}





