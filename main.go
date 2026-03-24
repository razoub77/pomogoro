package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/razoub77/pomogoro/timer"
)

const (
	defaultWorkDuration = 25
	defaultRestDuration = 5
)

func main() {
	workDuration := defaultWorkDuration
	restDuration := defaultRestDuration

	if len(os.Args) != 3 {
		fmt.Printf("Using default values of %d minutes working and %d minutes resting\n",
			defaultWorkDuration, defaultRestDuration)
	} else {
		wdur, err1 := strconv.Atoi(os.Args[1])
		rdur, err2 := strconv.Atoi(os.Args[2])
		if err1 != nil || err2 != nil {
			fmt.Printf("Problem setting interval values, using default values of %d minutes working and %d minutes resting\n",
				defaultWorkDuration, defaultRestDuration)
		} else {
			workDuration = wdur
			restDuration = rdur
		}
	}

	t := timer.Timer{
		Start:        time.Now(),
		InWorkMode:   true,
		WorkDuration: workDuration * 60,
		RestDuration: restDuration * 60,
	}

	prevElapsed := 0

	for {
		elapsed := t.GetElapsedTime()
		if elapsed != prevElapsed {
			t.PrintTimeRemaining(elapsed)
			prevElapsed = elapsed
			if t.ShouldSwitchMode(elapsed) {
				t.Alert()
				t.SwitchMode()
			}
		}
	}
}
