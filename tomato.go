package main

import (
	"fmt"
	"os"
	"time"
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

var usage = `tomato usage:
  tomato 25s
  tomato 25m
  tomato 1h20m20s
`

var (
	timer    *time.Timer
	ticker   *time.Ticker
	queues   chan termbox.Event
	exitCode int
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

func tbinit() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func clear() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
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

// fix: 1h25m25s
func format(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	if h < 1 {
		return fmt.Sprintf("%d:%02d", m, s)
	} else {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
}

type Font map[rune][]string

func toText(str string) Text {
	text := make(Text, 0)
	for _, runeValue := range str {
		text = append(text, defaultFont[runeValue])
	}
	return text
}

// Draw the moment, Text is like "00:25"
func draw(startX, startY int, t Text) {
	clear()
	x, y := startX, startY
	for _, s := range t {
		for _, line := range s {
			for _, ch := range line {
				termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
				x++
			}
			x = startX
			y++
		}
		startX += s.width()
		x, y = startX, startY
	}
	flush()
}

func countdown(timeleft time.Duration) {
	w, h := termbox.Size()
	str := format(timeleft)
	text := toText(str)
	startX, startY := w/2-text.width()/2, h/2-text.height()/2

	start(timeleft)
	draw(startX, startY, text)

loop:
	for {
		select {
		case ev := <-queues:
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc) {
				exitCode = 2
				break loop
			}
			if ev.Ch == 'P' || ev.Ch == 'p' {
				stop()
			}
			if ev.Ch == 'C' || ev.Ch == 'c' {
				start(timeleft)
			}
		case <-ticker.C:
			timeleft -= time.Duration(tick)
			str = format(timeleft)
			text = toText(str)
			draw(startX, startY, text)
		case <-timer.C:
			break loop
		}
	}
	termbox.Close()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
	}
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println("time format error", usage)
	}

	tbinit()

	queues = make(chan termbox.Event)
	go func() {
		for {
			queues <- termbox.PollEvent()
		}
	}()
	timeleft := duration
	// start a tamato clock
	countdown(timeleft)

	// transfer an integer number of units to a duration
	bt := time.Duration(5 * time.Second)
	// start to break between tomatoes
	breaktime(bt)
}
