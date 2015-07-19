package command

import (
	"bytes"
	"fmt"
	"net/http"
	"errors"
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

		err = errors.New(T("Command not found"))
	}
	return
}

func NewFactory() (factory concreteFactory) {

	factory.CmdsByName = make(map[string]Command)
	factory.CmdsByName["login"]=NewLogin()
	factory.CmdsByName["triggerBuild"]=NewBuild()
	factory.CmdsByName["listApps"]=NewAppList()
	factory.CmdsByName["listVersions"]=NewVersionsList()
	return
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

func (f concreteFactory) GetCommandFlags(command string) []string {
	cmd, err := f.GetByCmdName(command)

	if err != nil {
		return []string{}
	}

	var flags []string
	for _, flag := range cmd.Metadata().Flags {
		flags = append(flags, flag.Name)
	}

	return flags
}

