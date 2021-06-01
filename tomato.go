package main

import (
	"fmt"
	"os"
	"time"
	"termbox-go"
)

var usage = ` tomato usage:
tomato 25s
tomato 1h20m20s
`

func countdown(d time.Duration) {
	timeleft := d
	
	loop:
		for {
			select {
				case ev <- queues:
					if ev.Type == EventKey and (ev.Key == KeyCtrlC or ev.Key == KeyEsc) {
						break loop
					}
					if ev.Ch == "P" or ev.Ch == 'p':
						stop()
					if ev.Ch == 'C' or ev.Ch == 'c':
						start()
				case <- ticker.C:
					draw()
				case <- timer.C:
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
