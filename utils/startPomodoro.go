package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"
)

var (
	counter int = 0
)

func StartPomodoro(duration, pause time.Duration) {

	flagColorSent := SendColor()
	loadingBarColor := SetLoadingBarColor(flagColorSent)
	color.New(color.FgBlack, loadingBarColor).Println("Pomodoro started!")

	loadingBarWidth := 50
	bar := ""

	for {
		counter++
		for remaining := duration; remaining >= 0; remaining -= time.Second {
			ClearScreen()
			fmt.Printf("\nTime to go to work (%d/4)\n", counter)

			progress := int((float64(duration-remaining) / float64(duration)) * float64(loadingBarWidth))
			bar = color.New(color.FgBlack, loadingBarColor).Sprintf("%-*s", loadingBarWidth, strings.Repeat("█", progress))

			fmt.Printf("\r%*s\rTime remaining: %v [%s]", len("Time remaining: 10m10s"), "", remaining.Truncate(time.Second), bar)
			time.Sleep(time.Second)
		}

		if counter%4 != 0 {
			fmt.Printf("\nTime to take a break! (%v/4) \n", counter)
			for remaining := pause; remaining >= 0; remaining -= time.Second {
				progress := int((float64(pause-remaining) / float64(pause)) * float64(loadingBarWidth))
				bar = color.New(color.FgMagenta, loadingBarColor).Sprintf("%-*s", loadingBarWidth, strings.Repeat("█", progress))
				fmt.Printf("\r%*s\rTime remaining: %v [%s]", len("Time remaining: 10m10s"), "", remaining.Truncate(time.Second), bar)
				time.Sleep(time.Second)
			}
		} else {
			fmt.Printf("\nTime to take a long break! (%v/4)\n", counter)
			longBreak := time.Duration(pause * 3)
			for remaining := longBreak; remaining >= 0; remaining -= time.Second {
				progress := int((float64(longBreak-remaining) / float64(longBreak)) * float64(loadingBarWidth))
				bar := color.New(color.FgYellow, loadingBarColor).Sprintf("%-*s", loadingBarWidth, strings.Repeat("█", progress))
				fmt.Printf("\r%*s\rTime remaining: %v [%s]", len("Time remaining: 10m10s"), "", remaining.Truncate(time.Second), bar)
				time.Sleep(time.Second)
			}
			fmt.Printf("\nThe cycle starts again (%v/4) \n", counter)
			counter = 0
		}
	}
}

func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	cmd.Run()
}
