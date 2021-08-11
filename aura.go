package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/thatisuday/commando"
)

const (
	NAME    string = "aura"
	VERSION string = "1.3.0"
)

func downloadFile(filepath, url string) error {
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

	bar := progressbar.DefaultBytes(resp.ContentLength, filepath)

	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
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
		AddFlag("filename,f", "File name where your file will be executed", commando.String, ":detect").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

			url := args["url"].Value
			comp := strings.Split(url, "/")
			filename, e := flags["filename"].GetString()
			if e != nil {
				filename = comp[len(comp)-1]
			}

			if flags["filename"].Value.(string) == ":detect" {
				filename = comp[len(comp)-1]
				err := downloadFile(filename, url)
				if err != nil {
					fmt.Println("Unable to download file:", err)
				}
				fmt.Println("Downloaded: " + url)

			} else {
				err := downloadFile(filename, url)
				if err != nil {
					fmt.Println("Unable to download file:", err)
				}

				fmt.Println("Downloaded: " + url)
			}
		})

	commando.Parse(nil)
}
