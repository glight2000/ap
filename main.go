package main

import (
	"ap/action"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"runtime"
)

var (
	GOOS   = runtime.GOOS
	GOROOT = os.Getenv("GOROOT")
	GOPATH = os.Getenv("GOPATH")
	GOBIN  = os.Getenv("GOBIN")
	USER   = os.Getenv("USER")
	HOME   = os.Getenv("HOME")
)

/**
 have not tested on other os
 */
func main() {
	app := cli.NewApp()
	app.Name = "Apport"
	app.Usage = "[]"
	app.Action = func(c *cli.Context) {
		args := c.Args()
		if len(args) > 0 {
			switch args[0] {
			case "init":
				doInit()
			case "to":
				err := action.ApportGoto(c.Args()[1:])
				if err!=nil{
					log.Fatal(err)
				}
			case "t":
				err := action.ApportGoto(c.Args()[1:])
				if err!=nil{
					log.Fatal(err)
				}
			default:
				err := action.ApportRun(c.Args())
				if err!=nil{
					log.Fatal(err)
				}
			}
		}
	}
	app.Run(os.Args)
}

func doInit() {

	fmt.Println("GOOS :", GOOS)
	fmt.Println("GOROOT :", GOROOT)
	fmt.Println("GOPATH :", GOPATH)
	fmt.Println("GOBIN :", GOBIN)
	fmt.Println("USER :", USER)
	fmt.Println("HOME :", HOME)

	if len(GOBIN) < 0 {
		log.Fatal("Envirment path 'GOBIN' required!")
		return
	}

	//create default config file apport.cfg
	filePath := GOBIN + "ap.cfg"
	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)

	if err != nil {
		log.Fatal("create file ", filePath, "failed!")
		return
	}
	defer file.Close()

	if _, err := file.Write([]byte(CONF_TEMPLATE_WINDOWS)); err != nil {
		log.Fatal("write file ", filePath, "failed!")
		return
	}

	//install ar.go at.go to GOBIN
	switch GOOS {
	case "windows":

	case "linux":

	}
}

const (
	CONF_TEMPLATE_WINDOWS = `{
    "to": [
        {
            "alias": ["os"],
            "target": "c:/windows/"
        },
        {
            "alias": ["p01","css"],
            "target": "d:/dev/workspace/mvcprojectA/public/css/"
        }
    ],
    "customize": [
        {
            "shortcut": "open",
            "targets": [
                {
                    "app": "d:/dev/Sublime Text 3/sublime_text.exe",
                    "argument_filter": ".*\\.(txt|cfg|go)",
                    "ext_arguments":[],
                    "isArgumentsInherit":true
                },
                {
                    "app": "c:/windows/notepad.exe",
                    "argument_filter": ".*\\.(bat)",
                    "ext_arguments":[],
                    "isArgumentsInherit":true
                },
                {
                    "app": "c:/windows/explorer.exe",
                    "argument_filter": ".*\\.*|\\|.*/.*|/|\\.|",
                    "ext_arguments":[],
                    "isArgumentsInherit":true
                }
            ]
        },
        {
            "shortcut": "idea",
            "targets": [
                {
                    "app": "D:/dev/JetBrains/IntelliJ IDEA Community Edition 15.0.2/bin/idea.exe",
                    "argument_filter": "",
                    "ext_arguments":[],
                    "isArgumentsInherit":true
                }
            ]
        },
        {
            "shortcut": "www",
            "targets": [
                {
                    "app": "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe",
                    "argument_filter": "bd",
                    "ext_arguments":["http://www.baidu.com"],
                    "isArgumentsInherit":false
                },
                {
                    "app": "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe",
                    "argument_filter": "gg",
                    "ext_arguments":["https://www.google.com"],
                    "isArgumentsInherit":false
                }
            ]
        },
        {
            "shortcut": "game",
            "targets": [
                {
                    "app": "E:/Game/Warcraft III/war3.exe",
                    "argument_filter": "war3",
                    "ext_arguments":["-window"],
                    "isArgumentsInherit":false
                }
            ]
        }
    ]
}`
)
