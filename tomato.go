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

func format(d time.Duration) string {
	h := fmt.Sprintf("%02.f", d.Hours())
	m := fmt.Sprintf("%02.f", d.Minutes())
	s := fmt.Sprintf("%02.f", d.Seconds())
	if h == "0" {
		return fmt.Sprintf("%v:%v:%v", h, m, s)
	} else {
		str := fmt.Sprintf("%v:%v", m, s)
		return str
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

// 画此刻的Text
func draw(startX, startY int, t Text) {
	x, y := startX, startY
	for _, s := range t {
		for _, line := range s {
			for _, ch := range line {
				termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
				x += 1
			}
			x = startX
			y += 1
		}
		startX += s.width()
		x = startX
		y = startY
	}
	flush()
}

func countdown(d time.Duration) {
	w, h := termbox.Size()
	clear()
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
			timeleft -= 1 * time.Second
			str = format(timeleft)
			text = toText(str)
			fmt.Println(text)
			// draw(startX, startY, text)
		case <-timer.C:
			break loop
		}

	}
	termbox.Close()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
	}
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println("time format error", usage)
	}
	termbox.Init()
	go func() {
		queues <- termbox.PollEvent()
	}()

	countdown(duration)
}

[
	[ ██████╗  ██╔═████╗ ██║██╔██║ ████╔╝██║ ╚██████╔╝  ╚═════╝ ]
	[ ██████╗  ██╔═████╗ ██║██╔██║ ████╔╝██║ ╚██████╔╝  ╚═════╝ ]
	 [    ██╗ ╚═╝ ██╗ ╚═╝    ]
	  [ ██╗ ███║ ╚██║  ██║  ██║  ╚═╝]
	  [██████╗  ╚════██╗  █████╔╝  ╚═══██╗ ██████╔╝ ╚═════╝ ]
	  ]

[
	[ ██████╗  ██╔═████╗ ██║██╔██║ ████╔╝██║ ╚██████╔╝  ╚═════╝ ]
	[ ██████╗  ██╔═████╗ ██║██╔██║ ████╔╝██║ ╚██████╔╝  ╚═════╝ ]
	[    ██╗ ╚═╝ ██╗ ╚═╝    ]
	 [ ██╗ ███║ ╚██║  ██║  ██║  ╚═╝]
	[██████╗  ╚════██╗  █████╔╝ ██╔═══╝  ███████╗ ╚══════╝]
	]

[
	[ ██████╗  ██╔═████╗ ██║██╔██║ ████╔╝██║ ╚██████╔╝  ╚═════╝ ]
	[ ██████╗  ██╔═████╗ ██║██╔██║ ████╔╝██║ ╚██████╔╝  ╚═════╝ ]
	 [    ██╗ ╚═╝ ██╗ ╚═╝    ]
	 [ ██╗ ███║ ╚██║  ██║  ██║  ╚═╝]
	  [ ██╗ ███║ ╚██║  ██║  ██║  ╚═╝]]
