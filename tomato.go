package main

import (
	"fmt"
	"os"
	"time"
	"unicode/utf8"
	"github.com/nsf/termbox-go"
)

var usage = ` tomato usage:
tomato 25s
tomato 25m
tomato 1h20m20s
`

var (
	timer  *time.Timer
	ticker *time.Ticker
	queues chan termbox.Event 
)

const (
	tick = time.Second
)

type Symbol []string

func (s Symbol) width() int {
	return utf8.RuneCountInString(s[0])
}

func (s Symbol) height() int {
	return len(s)
}

type Text []Symbol

func (t Text) width() (w int) {
	w = 0
	for _, s := range t {
		// w += utf8.RuneCountInString(s[0])
		w += s.width()
	}
	return 
}

func (t Text) height() int {
	return len(t[0])
}

func start(d time.Duration) {
	timer = time.NewTimer(d)
	ticker = time.NewTicker(tick)
}

func stop() {
	timer.Stop()
	ticker.Stop()
}

func clear() {
	err := termbox.Clear(termbox.ColorBlue, termbox.ColorBlue)
	if err != nil {
		panic(err)
	}
}

func flush() {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func format(d time.Duration) string {
	h := fmt.Sprintf("%.0f", d.Hours())
	m := fmt.Sprintf("%.0f", d.Minutes())
	s := fmt.Sprintf("%.0f", d.Seconds())
	if h == "0" {
		str := fmt.Sprintf("%2.f:%2.f:%2.f", h, m, s)
	} else {
		str := fmt.Sprintf("%2.f:%2.f", m, s)
	}
	return str
}

// 画此刻的Text
func draw(startX, startY int, t Text) {
	x, y := startX, startY
	for _, s := range t {
		for _, line := range s {
			for _, ch := range line {
				termbox.SetCell(x, y, ch, termbox.ColorBlue, termbox.ColorBlue)
			}
		} 
	}
	flush()
}

func countdown(d time.Duration) {
	w, h := termbox.Size()
	str := format(d)
	text := toText(str)
	startX, startY := w / 2 - text.width(), h / 2 - text.height() / 2
	timeleft := d
	start(timeleft)
	draw(startX, startY)


loop:
	for {
		select {
		case ev := <-queues:
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc) {
				break loop
			}
			if ev.Ch == 'P' || ev.Ch == 'p' {
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
		fmt.Sprintln(usage)
	}
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println("time format error", usage)
	}
	go func() {
		queues <- termbox.PollEvent()
	}()
	
	countdown(duration)
}
