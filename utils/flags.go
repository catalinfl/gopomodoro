package utils

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	duration = time.Duration(25 * time.Minute)
	pause    = time.Duration(5 * time.Minute)
	color    = "No color"
)

func Flags() {
	helpFlag := flag.Bool("help", false, "Show help")
	setFlag := flag.Int("set", 0, "Set duration of the pomodoro")
	// flag.StringVar(&color, "color", "red", "Set color of the pomodoro")

	flag.Parse()

	if *helpFlag {
		fmt.Println("use: gopomodoro [options]")
		fmt.Println("- help = show help cmds")
		fmt.Println("- set = set duration of the pomodoro, pause is duration / 5")
		os.Exit(0)
	}

	if *setFlag >= 5 {
		duration = time.Duration(*setFlag) * time.Minute
		pause = time.Duration(*setFlag/5) * time.Minute
		return
	} else if *setFlag < 5 {
		fmt.Println("Duration must be greater than 5 minutes. Using default values.")
	}

	fmt.Println("Invalid command")
}

func SendTime() (time.Duration, time.Duration, error) {
	return duration, pause, nil
}

// func SendColor() string {
// 	return color
// }
