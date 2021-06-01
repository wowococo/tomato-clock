package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

var usage = ` tomato usage:
tomato 25s
tomato 1h20m20s
`

var (
	timer  *time.Timer
	ticker *time.Ticker
)

const (
	tick = time.Second
)

func start(d time.Duration) {
	timer = time.NewTimer(d)
	ticker = time.NewTicker(tick)
}

func stop() {
	timer.Stop()
	ticker.Stop()
}

func draw() {

}

func countdown(d time.Duration) {
	timeleft := d

loop:
	for {
		select {
		case ev := <-queues:
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc) {
				break loop
			}
			if ev.Ch == "P" || ev.Ch == 'p' {
				stop()
			}
			if ev.Ch == 'C' || ev.Ch == 'c' {
				start(d)
			}
		case <-ticker.C:
			draw()
		case <-timer.C:
			break loop
		}

	}
}

func main() {
	if len(os.Args) != 2 {
		os.Stderr(usage)
	}
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println("time format error", usage)
	}

	countdown(duration)
}
