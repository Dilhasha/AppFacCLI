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

type AppList struct {
	//to be added
}

func NewAppList() (cmd AppList) {
	return
}

func (applist AppList)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"getApplicationsOfUser",
		Description : "Lists applications of a user",
		ShortName : "la",
		Usage:"list apps",
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/application/get/ajax/list.jag",
		SkipFlagParsing:false,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "-u", Usage: "userName"},
			cli.StringFlag{Name: "-c", Usage: "cookie"},
		},
	}

}


func(applist AppList) Run(c CommandConfigs)(bool,string){
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
				fmt.Println("You have ", len(apps)," applications. Details of applications are as follows.\n")
				totals := tm.NewTable(0, 10, 5, ' ', 0)
				fmt.Fprintf(totals, "Name\tKey\tType\tDescription\n")
				fmt.Fprintf(totals, "----\t---\t----\t-----------\n")
				for _, app := range apps {
					fmt.Fprintf(totals, "%s\t%s\t%s\t%s\n", app.Name,app.Key,app.Type,app.Description)
				}
				tm.Println(totals)
				tm.Flush()
			}



		}
	}
	return true,c.Cookie
}
