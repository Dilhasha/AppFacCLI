package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/Dilhasha/AppFacCLI/cli/formats"
	"github.com/codegangsta/cli"
	tm "github.com/buger/goterm"
)

type AppInfo struct {
	//to be added
}

func NewAppInfo() (cmd AppInfo) {
	return
}

func (appInfo AppInfo)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"getAppInfo",
		Description : "get information of an application",
		ShortName : "ai",
		Usage:"get app info",
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/application/get/ajax/list.jag",
		SkipFlagParsing:false,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "-a", Usage: "applicationKey"},
			cli.StringFlag{Name: "-c", Usage: "cookie"},
		},
	}
}


func(appInfo AppInfo) Run(c CommandConfigs)(bool,string){
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
			println("has error")
		}
		var app formats.AppFormat
		err = json.Unmarshal([]byte(bodyStr), &app)
		if(err ==nil){
				fmt.Println("\nDetails of application are as follows.\n")
				totals := tm.NewTable(0, 10, 5, ' ', 0)
				fmt.Fprintf(totals, "Name\tKey\tType\tRepositoryType\tOwner\n")
				fmt.Fprintf(totals, "-------\t------\t-----\t---------\t----------------\n")
				fmt.Fprintf(totals, "%s\t%s\t%s\t%s\t%s\n", app.Name,app.Key,app.Type,app.RepositoryType,app.Owner)
				tm.Println(totals)
				tm.Flush()
		}


	}
	return true,c.Cookie
}
