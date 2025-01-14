package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"

	systray "github.com/getlantern/systray"
)

const (
	TMP_FILE        = "/tmp/wakatime"
	UPDATE_INTERVAL = 15 * time.Minute
	ICON_PATH       = "/home/celso/projects/cmdtray/icons/favicon.ico"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	readIcon()
	quitButton := systray.AddMenuItem("Quit", "quit")

	runCommandAndUpdate()
	ticker := time.NewTicker(UPDATE_INTERVAL)

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
		slog.Error("Error running command", "err", err.Error())
		return
	}

	out := string(res)
	slog.Info("Wakatime:", "out", out)
	writeFile(out)

	systray.SetTitle(
		strings.TrimSpace(out),
	)
}

func writeFile(s string) {
	err := os.WriteFile(TMP_FILE, []byte(s), 0777)
	if err != nil {
		slog.Error("Error writing file", "err", err.Error())
		return
	}
	slog.Info("Written to file", "file", TMP_FILE)
}

func readIcon() {
	buf, err := os.ReadFile(ICON_PATH)
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	systray.SetIcon(buf)
}

func onExit() {}
