package utils

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	duration  = time.Duration(25 * time.Minute)
	pause     = time.Duration(5 * time.Minute)
	flagColor = "No color"
)

func Flags() {
	helpFlag := flag.Bool("help", false, "Show help")
	setFlag := flag.Int("set", 0, "Set duration of the pomodoro")
	flag.StringVar(&flagColor, "color", "cyan", "Set color of the pomodoro")

	flag.Parse()

	if *helpFlag {
		fmt.Println("use: gopomodoro [options]")
		fmt.Println("-help = show help cmds")
		fmt.Println("-set <int> = set duration of the pomodoro, pause is duration / 5")
		fmt.Println("-color [red, green, yellow, blue, magenta, cyan, white] = set color of the pomodoro")
		fmt.Println("example: gopomodoro -set 35 -color green")
		os.Exit(0)
	}

	if *setFlag >= 5 {
		duration = time.Duration(*setFlag) * time.Minute
		pause = time.Duration(*setFlag/5) * time.Minute
		return
	} else if *setFlag == 0 {
		return
	} else if *setFlag < 5 && *setFlag != 0 {
		fmt.Println("Duration must be greater than 5 minutes. Using default values.")
		return
	}

	fmt.Println("Invalid flag. Using default values.")
}

func SendTime() (time.Duration, time.Duration, error) {
	return duration, pause, nil
}

func SendColor() string {
	return flagColor
}
