package utils

import (
	"fmt"
	"time"
)

var (
	counter int = 0
)

func StartPomodoro(duration, pause time.Duration) {
	for {
		for remaining := duration; remaining >= 0; remaining -= time.Second {
			fmt.Printf("\r%*s\rTime remaining: %v", len("Time remaining: 10m10s"), "", remaining.Truncate(time.Second))
			time.Sleep(time.Second)
		}

		counter++

		if counter%4 != 0 {
			fmt.Println("\nTime to take a break!")
			for remaining := pause; remaining >= 0; remaining -= time.Second {
				fmt.Printf("\r%*s\rTime remaining: %v", len("Time remaining: 10m10s"), "", remaining.Truncate(time.Second))
				time.Sleep(time.Second)
			}
		} else {
			fmt.Println("\nTime to take a long break!")
			longBreak := time.Duration(pause * 3)
			for remaining := longBreak; remaining >= 0; remaining -= time.Second {
				fmt.Printf("\r%*s\rTime remaining: %v", len("Time remaining: 10m10s"), "", remaining.Truncate(time.Second))
				time.Sleep(time.Second)
			}
			fmt.Printf("\nThe cycle starts again (%v/4) \n", counter)
			counter = 0
		}

		fmt.Printf("\nTime to go to work (%d/4)\n", counter)
	}

}
