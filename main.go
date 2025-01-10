package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	systray "github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	readIcon()
	quitButton := systray.AddMenuItem("Quit", "quit")

	runCommandAndUpdate()
	ticker := time.NewTicker(2 * time.Minute)

	for {
		select {
		case <-ticker.C:
			runCommandAndUpdate()
		case <-quitButton.ClickedCh:
			systray.Quit()
		}
	}
}

func runCommandAndUpdate() {
	cmd := exec.Command("wakatime-cli", "--today")

	res, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running command", err.Error())
		return
	}

	systray.SetTitle(
		strings.TrimSpace(string(res)),
	)
}

func readIcon() {
	buf, err := os.ReadFile("/home/celso/projects/cmdtray/icons/favicon.ico")
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	systray.SetIcon(buf)
}

func onExit() {}
