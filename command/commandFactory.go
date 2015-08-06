package command

import (
	"bytes"
	"fmt"
	"net/http"
	"errors"
	"github.com/codegangsta/cli"
)

type concreteFactory struct {
	CmdsByName map[string]Command
}

func (f concreteFactory) GetByCmdName(cmdName string) (cmd Command, err error) {
	cmd, found := f.CmdsByName[cmdName]
	if !found {
		for _, c := range f.CmdsByName {
			if c.Metadata().ShortName == cmdName {
				return c, nil
			}
		}

		err = errors.New("Command not found")
	}
	return
}

func NewFactory() (factory concreteFactory) {

	factory.CmdsByName = make(map[string]Command)
	factory.CmdsByName["login"]=NewLogin()
	factory.CmdsByName["triggerBuild"]=NewBuild()
	factory.CmdsByName["listApps"]=NewAppList()
	factory.CmdsByName["listVersions"]=NewVersionsList()
	factory.CmdsByName["createApp"]=NewAppCreation()
	factory.CmdsByName["exit"]=NewExit()
	return
}

func (f concreteFactory) GetCommandFlags(cmd Command) []string {
	var flags []string
	for _, flag := range cmd.Metadata().Flags {
		switch t := flag.(type) {
		default:
		case cli.StringFlag:
			flags = append(flags, t.Name)
		}
	}

	return flags
}


func (f concreteFactory) GetCommandConfigs(cmd Command,flagVals []string) CommandConfigs {
	var buffer bytes.Buffer
	flags:=cmd.Metadata().Flags
	var cookie string
	buffer.WriteString("action="+cmd.Metadata().Name)

	for n := 0; n < len(flags); n++ {

		if flag, ok := flags[n].(cli.StringFlag); ok {
			if(flag.Usage=="cookie"){
				cookie=flagVals[n]
			}else{
				buffer.WriteString("&"+flag.Usage+"=")
				buffer.WriteString(flagVals[n])
			}
		}


	}
	s := buffer.String()
	return CommandConfigs{
		Url:cmd.Metadata().Url,
		Query:s,
		Cookie:cookie,
	}
}


func (c CommandConfigs) Run() (*http.Response){
	fmt.Println("URL:>", c.Url)
	var jsonStr = []byte(c.Query)
	req, err := http.NewRequest("POST", c.Url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("Cookie", c.Cookie)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	/*fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header.Get("Content-Type"))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))*/

	return resp
}



