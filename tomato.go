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
	timer  *time.Timer
	ticker *time.Ticker
	queues chan termbox.Event
)

const (
	tick = time.Second
)

type Symbol []string

func (s Symbol) width() int {
	fmt.Println(s)
	return utf8.RuneCountInString(s[0])
}

func (s Symbol) height() int {
	return len(s)
}

type Text []Symbol

func (t Text) width() (w int) {
	w = 0
	for _, s := range t {
		fmt.Println(t)
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
	h := fmt.Sprintf("%2.f", d.Hours())
	m := fmt.Sprintf("%2.f", d.Minutes())
	s := fmt.Sprintf("%2.f", d.Seconds())
	fmt.Println(h, m, s)
	if h == "00" {
		return fmt.Sprintf("%d:%d:%d", h, m, s)
	} else {
		return fmt.Sprintf("%d:%d", m, s)
	}
}

type Font map[rune][]string

func toText(str string) Text {
	fmt.Println(str)
	text := make(Text, 0)
	for _, runeValue := range str {
		text = append(text, defaultFont[runeValue])
	}
	return text
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
	startX, startY := w/2-text.width()/2, h/2-text.height()/2
	timeleft := d
	start(timeleft)
	draw(startX, startY, text)

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
			timeleft -= 1
			str = format(timeleft)
			text = toText(str)
			draw(startX, startY, text)
		case <-timer.C:
			break loop
		}

	}
}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 2 {
		fmt.Println(usage)
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
