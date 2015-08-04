package command

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"github.com/codegangsta/cli"

)

type Login struct {
	//to be added
}

func NewLogin() (cmd Login) {
	return
}

func (login Login)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"login",
		Description : "Login to app factory",
		ShortName : "l",
		Usage:"login",
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/user/login/ajax/login.jag",
		SkipFlagParsing:true,
		Flags: []cli.Flag{
			cli.StringFlag{Name:"-u",Usage:"userName"},
			cli.StringFlag{Name: "-p", Usage: "password"},
		},
	}
}



func(login Login) Run(c CommandConfigs)(bool,string){
	var resp *http.Response
	var bodyStr string
	resp=c.Run()
	defer resp.Body.Close()

	if(resp.Status=="200 OK"){
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr=string(body)
		if(strings.Contains(bodyStr, "true")){
			fmt.Println("You have Successfully logged in.")
			cookie:=strings.Split(resp.Header.Get("Set-Cookie"),";")
			c.Cookie=cookie[0]
			fmt.Println("Cookie for the session is:",cookie[0])
		}else{
			fmt.Println("Authorization failed. Please try again!")
		}
	}
	return true,c.Cookie
}

