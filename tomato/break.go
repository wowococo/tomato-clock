package tomato

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func breaktime(d time.Duration, tomatoID int64) {
	echo()
	tbinit()
	countdown(d, tomatoID)

}

func echo() {
	err := ui.Init()
	if err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	p := widgets.NewParagraph()

	p.Text = "TIME TO BREAK ^_^ "

	p.SetRect(0, 0, 25, 5)

	ui.Render(p)
	time.Sleep(2 * time.Second)

}
