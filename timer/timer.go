package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	Start        time.Time
	InWorkMode   bool
	WorkDuration int
	RestDuration int
}

func (t Timer) GetElapsedTime() int {
	return int(time.Since(t.Start).Seconds())
}

func (t *Timer) SwitchMode() {
	t.Start = time.Now()
	t.InWorkMode = !t.InWorkMode
}

func (t Timer) ShouldSwitchMode(elapsed int) bool {
	return elapsed == t.getDuration()
}

func (t Timer) Alert() {
	msg := "\nTake a break"
	if !t.InWorkMode {
		msg = "\nBack to work"
	}
	fmt.Println(msg)
}

func (t Timer) getDuration() int {
	duration := t.WorkDuration
	if !t.InWorkMode {
		duration = t.RestDuration
	}
	return duration
}

func (t Timer) getMode() string {
	mode := "Working"
	if !t.InWorkMode {
		mode = "Resting"
	}
	return mode
}

func (t Timer) PrintTimeRemaining(elapsed int) {
	timeRemaining := t.getDuration() - elapsed
	min := timeRemaining / 60
	sec := timeRemaining - min*60
	fmt.Printf("\r%v: %02dm:%02ds", t.getMode(), min, sec)
}
