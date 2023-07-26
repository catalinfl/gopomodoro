package main

import (
	"fmt"
	"os"

	"github.com/catalinfl/gopomodoro/utils"
)

func main() {

	utils.ClearScreen()
	utils.Flags()
	duration, pause, err := utils.SendTime()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Starting pomodoro with %v minutes of work and %v minutes of pause...\n", duration.Minutes(), pause.Minutes())

	utils.StartPomodoro(duration, pause)
}
