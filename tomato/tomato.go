package tomato

import (
	"flag"
	"fmt"
	"os"
	"time"
	"tomato-clock/sqliteopt"
	"tomato-clock/stats"
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

const (
	Running  = 0
	Finished = 1
	Pause    = 2
	Dropout  = 3
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
func draw(startX, startY int, t Text, bk bool) {
	clear()
	x, y := startX, startY
	var fg, bg termbox.Attribute
	if bk {
		fg = termbox.ColorGreen
	} else {
		fg = termbox.ColorDefault
	}
	bg = termbox.ColorDefault
	for _, s := range t {
		for _, line := range s {
			for _, ch := range line {
				termbox.SetCell(x, y, ch, fg, bg)
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

func countdown(timeleft time.Duration, tomatoID int64, bk bool) {
	// tomato initial duration
	d := timeleft
	w, h := termbox.Size()
	str := format(timeleft)
	text := toText(str)
	startX, startY := w/2-text.width()/2, h/2-text.height()/2

	start(timeleft)
	draw(startX, startY, text, bk)

	// Execute only when you are focused on your time
	notbk := func(bk bool, tomatoID int64, timeleft, d time.Duration, status int8) {
		if !bk {
			sqliteopt.PutTomato(tomatoID, timeleft, d, status)
		}
	}

loop:
	for {
		select {
		case ev := <-queues:
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc) {
				// sqliteopt.PutTomato(tomatoID, timeleft, d, Dropout)
				notbk(bk, tomatoID, timeleft, d, Dropout)
				exitCode = 2
				break loop
			}
			if ev.Ch == 'P' || ev.Ch == 'p' {
				stop()
				// sqliteopt.PutTomato(tomatoID, timeleft, d, Pause)
				notbk(bk, tomatoID, timeleft, d, Pause)
			}
			if ev.Ch == 'C' || ev.Ch == 'c' {
				notbk(bk, tomatoID, timeleft, d, Running)
				start(timeleft)
			}
		case <-ticker.C:
			timeleft -= time.Duration(tick)
			str = format(timeleft)
			text = toText(str)
			draw(startX, startY, text, bk)
		case <-timer.C:
			// sqliteopt.PutTomato(tomatoID, timeleft, d, Finished)
			notbk(bk, tomatoID, 0, d, Finished)
			break loop
		}
	}

	termbox.Close()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

const (
	ENDTASK = "endtask"
)

func Tomato() {
	var d = flag.Duration("d", 25*time.Minute, "tomato clock duration")
	var bt = flag.Duration("bt", 5*time.Minute, "break time duration")
	var t = flag.String("t", "Unnamed", "task name")
	// bool type don't like string type need to input true or false, if you only inout "-chart",
	// then it is true, even you input "-chart false", it also be true.
	var chart = flag.Bool("chart", false, "show report form, metrics and linechart")
	var endtask = flag.String("endtask", "", "mark a task finished")
	flag.Parse()

	// why this three lines can't print
	// fmt.Println("hhhhhh", len(os.Args))
	// if len(os.Args) == 1 {
	// 	flag.PrintDefaults()
	// }

	duration := *d

	tbinit()

	queues = make(chan termbox.Event)
	go func() {
		for {
			queues <- termbox.PollEvent()
		}
	}()

	var flagset = make(map[string]flag.Value)
	flag.Visit(func(f *flag.Flag) {
		flagset[f.Name] = f.Value
	})

	if len(flagset) == 0 {
		flag.PrintDefaults()
	}

	if _, ok := flagset["d"]; ok {
		timeleft := duration
		// start a tamato clock
		taskID, ok := sqliteopt.GetTask(*t)
		if ok {
			sqliteopt.PutTask("", taskID)
		} else {
			if *t != "" {
				taskID = sqliteopt.PostTask(*t, Running)
			} else {
				os.Exit(2)
			}
		}

		tomatoID := sqliteopt.PostTomato(taskID, duration, Running)
		countdown(timeleft, tomatoID, false)

		// start to break between tomatoes
		breaktime(*bt, tomatoID)
	}

	if *endtask != "" {
		taskID, ok := sqliteopt.GetTask(*endtask)
		if ok {
			sqliteopt.PutTask(ENDTASK, Finished, taskID)
		}

	}

	// stats
	if *chart {
		stats.Draw()
	}
}
