package stats

import (
	"log"
	"tomato-clock/sqliteopt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/linechart"
)

func drawText() {
	p0 := widgets.NewParagraph()
	p0.Title = "总完成番茄数"
	p0.Text = sqliteopt.Query("tomato", "progress", "total")
	p0.TextStyle = ui.NewStyle(ui.ColorRed)
	p0.SetRect(0, 0, 15, 3)
	p0.Border = true

	p1 := widgets.NewParagraph()
	p1.Title = "本周完成番茄数"
	p1.Text = sqliteopt.Query("tomato", "progress", "thisweek")
	p1.TextStyle = ui.NewStyle(ui.ColorRed)
	p1.SetRect(15, 0, 31, 3)
	p1.Border = true

	p2 := widgets.NewParagraph()
	p2.Title = "今日完成番茄数"
	p2.Text = sqliteopt.Query("tomato", "progress", "today")
	p2.TextStyle = ui.NewStyle(ui.ColorRed)
	p2.SetRect(32, 0, 48, 3)
	p2.Border = true

	p3 := widgets.NewParagraph()
	p3.Title = "总专注时间"
	p3.Text = sqliteopt.Query("tomato", "timefocused", "total")
	p3.TextStyle = ui.NewStyle(ui.ColorRed)
	p3.SetRect(0, 3, 15, 6)
	p3.Border = true

	p4 := widgets.NewParagraph()
	p4.Title = "本周专注时间"
	p4.Text = sqliteopt.Query("tomato", "timefocused", "thisweek")
	p4.SetRect(15, 3, 31, 6)
	p4.TextStyle = ui.NewStyle(ui.ColorRed)
	p4.Border = true

	p5 := widgets.NewParagraph()
	p5.Title = "今日专注时间"
	p5.Text = sqliteopt.Query("tomato", "timefocused", "today")
	p5.TextStyle = ui.NewStyle(ui.ColorRed)
	p5.SetRect(32, 3, 48, 6)
	p5.Border = true

	p6 := widgets.NewParagraph()
	p6.Title = "总完成任务"
	p6.Text = sqliteopt.Query("task", "id", "total")
	p6.TextStyle = ui.NewStyle(ui.ColorRed)
	p6.SetRect(0, 6, 15, 9)
	p6.Border = true

	p7 := widgets.NewParagraph()
	p7.Title = "本周完成任务"
	p7.Text = sqliteopt.Query("task", "id", "thisweek")
	p7.TextStyle = ui.NewStyle(ui.ColorRed)
	p7.SetRect(15, 6, 31, 9)
	p7.Border = true

	p8 := widgets.NewParagraph()
	p8.Title = "今日完成任务"
	p8.Text = sqliteopt.Query("task", "id", "today")
	p8.TextStyle = ui.NewStyle(ui.ColorRed)
	p8.SetRect(32, 6, 48, 9)
	p8.Border = true

	ui.Render(p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func inputs() {

}

func playLineChart() {

}

func drawLine() {
	t, err := tcell.New()
	hdlerr(err)
	defer t.Close()

	lc, err := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorDefault)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorBlue)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorFuchsia)),
	)
	hdlerr(err)

	go playLineChart()

	c, err := container.New(
		t,
		container.BorderTitle("tomatoes in a week"),
		container.BorderColor(cell.ColorRed),
		container.PlaceWidget(lc),
	)
	hdlerr(err)

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == "q"|k.Key == "Q" {
			return
		}
	}

	_, err = termdash.NewController(t, c, termdash.KeyboardSubscriber(quitter))
	hdlerr(err)
}

func Draw() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	go drawText()
	go drawLine()

	uiEvents := ui.PollEvents()
	for {

		switch ev := <-uiEvents; ev.ID {
		case "q", "<C-c>":
			return // if break is not working
		}
	}
}

func hdlerr(err error) {
	if err != nil {
		panic(err)
	}
}
