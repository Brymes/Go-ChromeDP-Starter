package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func LaunchChrome() {
	hostOS := runtime.GOOS

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if hostOS == "windows" {

		// Check if File exists in Program Files or APPDATA
		if err := os.Mkdir("chromedata", os.ModePerm); err != nil {
			log.Fatal(err)
		}

		userdir := path + "/chromedata"
		userdir = fmt.Sprintf("--user-data-dir=%v", userdir)

		cmd := exec.Command(`start chrome.exe`, "--remote-debugging-port=9222", userdir)

		go func() {
			cmd.Run()
		}()

		return
	} else if hostOS == "linux" {
		log.Fatal("You need to Fix this")

	} else if hostOS == "darwin" {
		err = os.RemoveAll(path + "/utils/~")
		if err != nil {
			log.Println(err)
		}
		//cmd := exec.Command("/Applications/Google\\ Chrome.app/Contents/MacOS/Google\\ Chrome", " --remote-debugging-port=9222", "--user-data-dir=~/chromedevelop/")
		cmd := exec.Command(`/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`, "--remote-debugging-port=9222", "--user-data-dir=utils/~")

		go func() {
			cmd.Run()
		}()

		return
	} else {
		log.Println(hostOS)
		log.Fatal("Please fix Chrome opening for the OS above")
	}
}
