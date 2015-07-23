package command

import (
	"fmt"
	"bytes"
	"bufio"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"io/ioutil"
	"net/http"
	"github.com/codegangsta/cli"

)
//Unix specific
const (
	sttyArg0   = "/bin/stty"
	exec_cwdir = ""
)

// Tells the terminal to turn echo off.
var sttyArgvEOff []string = []string{"stty", "-echo"}

// Tells the terminal to turn echo on.
var sttyArgvEOn []string = []string{"stty", "echo"}

var ws syscall.WaitStatus = 0


type Login struct {
	//to be added
}

func NewLogin() (cmd Login) {
	return
}

func (login Login)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"Login",
		Description : "Login to app factory",
		ShortName : "l",
		Usage:"login",
		SkipFlagParsing:true,
		Flags: []cli.Flag{
			cli.StringFlag{Name:"u",Usage:"userName"},
			cli.StringFlag{Name: "p", Usage: "password"},
		},
	}
}

func (login Login)Configs(reqs CommandRequirements)(configs CommandConfigs){
	var buffer bytes.Buffer
	buffer.WriteString("action=login")
	if(reqs.UserName!=""){
		buffer.WriteString("&userName=")
		buffer.WriteString(reqs.UserName)
	}
	if(reqs.Password!=""){
		buffer.WriteString("&password=")
		buffer.WriteString(reqs.Password)

	}

	s := buffer.String()
	return CommandConfigs{
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/user/login/ajax/login.jag",
		Query:s,
		Cookie:reqs.Cookie,
	}
}

func (login Login) Requirements(args []string)(reqs CommandRequirements){
	if(login.Metadata().SkipFlagParsing) {
		var username string
		fmt.Println("username: ")
		fmt.Scanf("%s", &username)
		//fmt.Println("password:")
		//fmt.Scanf("%s", &password)
		reqs.UserName = username
		reqs.Password = login.AskForPassword("Password")
		//reqs.Password=password
	}else{
		reqs.UserName=args[0]
		reqs.Password=args[1]
	}
	reqs.Cookie = ""
	return
}

func(login Login) Run(c CommandConfigs){
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
			fmt.Println("Cookie for the session is:",cookie[0])
		}else{
			fmt.Println("Authorization failed. Please try again!")
		}
	}


}
/**Ask Password functionality for unix*/

func (login Login) AskForPassword(prompt string) (passwd string) {
	sig := make(chan os.Signal, 10)

	// Display the prompt.
	fmt.Println("")
	fmt.Printf(prompt+": ")

	// File descriptors for stdin, stdout, and stderr.
	fd := []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()}

	// Setup notifications of termination signals to channel sig, create a process to
	// watch for these signals so we can turn back on echo if need be.
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT,
		syscall.SIGTERM)
	defer signal.Stop(sig)

	go catchSignal(fd, sig)

	pid, err := echoOff(fd)
	defer echoOn(fd)
	if err != nil {
		return
	}

	passwd = readPassword(pid)

	// Carriage return after the user input.
	fmt.Println("")

	return
}

func readPassword(pid int) string {
	rd := bufio.NewReader(os.Stdin)
	syscall.Wait4(pid, &ws, 0, nil)

	line, err := rd.ReadString('\n')
	if err == nil {
		return strings.TrimSpace(line)
	}
	return ""
}


// catchSignal tries to catch SIGKILL, SIGQUIT and SIGINT so that we can turn terminal
// echo back on before the program ends.  Otherwise the user is left with echo off on
// their terminal.
func catchSignal(fd []uintptr, sig chan os.Signal) {
	select {
	case <-sig:
		echoOn(fd)
		os.Exit(2)
	}
}


func echoOff(fd []uintptr) (int, error) {
	pid, err := syscall.ForkExec(sttyArg0, sttyArgvEOff, &syscall.ProcAttr{Dir: exec_cwdir, Files: fd})

	if err != nil {
		//Removed error and replaced with nil
		return 0, nil
	}
	return pid, nil
}

// echoOn turns back on the terminal echo.
func echoOn(fd []uintptr) {
	// Turn on the terminal echo.
	pid, e := syscall.ForkExec(sttyArg0, sttyArgvEOn, &syscall.ProcAttr{Dir: exec_cwdir, Files: fd})

	if e == nil {
		syscall.Wait4(pid, &ws, 0, nil)
	}
}
