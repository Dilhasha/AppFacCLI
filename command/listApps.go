package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/Dilhasha/AppFacCLI/cli/formats"
	"github.com/codegangsta/cli"
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
				for _, app := range apps {
					fmt.Println("\nName:\t"+app.Name)
					fmt.Println("------------------------------------------")
					fmt.Println("Key:\t"+app.Key)
					fmt.Println("Type:\t"+app.Type)
					fmt.Println("Description:\t"+app.Description)
					fmt.Print("Users:\t")
					for i:=0;i<len(app.Users);i++ {
						fmt.Print(app.Users[i].UserName)
						if(i!=len(app.Users)-1){
							fmt.Print(",")
						}else{
							fmt.Println()
						}
					}
					fmt.Print("InProduction:\t",app.InProduction)
				}
			}



		}
	}
	return true,c.Cookie
}
