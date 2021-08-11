package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/thatisuday/commando"
)

const (
	NAME    string = "aura"
	VERSION string = "1.2.0"
)

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	fmt.Println(NAME, VERSION)
	commando.
		SetExecutableName(NAME).
		SetVersion(VERSION).
		SetDescription("aura is a command line utility for fetching content from the web. It is efficient and fast, being written in golang.")
	commando.
		Register(nil).
		AddArgument("url", "The URL to download.", "").
		// comp := strings.Split(args["url"].Value , "/")
		AddFlag("fn", "File name where your file will be executed", commando.String, "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// fileUrl := "https://raw.githubusercontent.com/MatMoul/archfi/master/archfi"
			comp := strings.Split(args["url"].Value, "/")
			if flags["fn"].Value.(string) == "" {
				dir := comp[len(comp)-1]
				err := DownloadFile(dir, args["url"].Value)
				if err != nil {
					panic(err)
				}
				fmt.Println("Downloaded: " + args["url"].Value)
			} else {
				dir := flags["fn"].Value.(string)
				err := DownloadFile(dir, args["url"].Value)
				if err != nil {
					panic(err)
				}
				fmt.Println("Downloaded: " + args["url"].Value)
			}
		},

		// err := DownloadFile(dir , args["url"].Value )

		)
	commando.Parse(nil)
}
