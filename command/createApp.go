package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/codegangsta/cli"
	"github.com/Dilhasha/AppFacCLI/cli/formats"
	"encoding/json"
)

type AppCreation struct {
	//to be added
}

func NewAppCreation() (cmd AppCreation) {
	return
}

func (appCreation AppCreation)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"createNewApplication",
		Description : "Creates a new application",
		ShortName : "ca",
		Usage:"create application",
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/application/add/ajax/add.jag",
		SkipFlagParsing:false,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "-u", Usage: "userName"},
			cli.StringFlag{Name: "-c", Usage: "cookie"},
			cli.StringFlag{Name: "-k", Usage: "applicationKey"},
			cli.StringFlag{Name: "-n", Usage: "applicationName"},
			cli.StringFlag{Name: "-d", Usage: "applicationDescription"},
			cli.StringFlag{Name: "-t", Usage: "applicationType"},
			cli.StringFlag{Name: "-r", Usage: "repositoryType"},
		},
	}
	
}

func(appCreation AppCreation) Run(c CommandConfigs) (bool,string){
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
		}
		var successMessage formats.SuccessFormat
		err1 := json.Unmarshal([]byte(bodyStr), &successMessage)
		if(err1 ==nil){
			fmt.Println(successMessage.Message)
		}
	}
	return true,c.Cookie
}
