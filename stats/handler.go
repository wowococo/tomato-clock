package stats

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func drawText() {
	p0 := widgets.NewParagraph()
	p0.Title = "总完成番茄数"
	p0.Text = "927.7"
	p0.TextStyle = ui.NewStyle(ui.ColorRed)
	p0.SetRect(0, 0, 15, 3)
	p0.Border = true

	p1 := widgets.NewParagraph()
	p1.Title = "本周完成番茄数"
	p1.Text = "9"
	p1.TextStyle = ui.NewStyle(ui.ColorRed)
	p1.SetRect(15, 0, 31, 3)
	p1.Border = true

	p2 := widgets.NewParagraph()
	p2.Title = "今日完成番茄数"
	p2.Text = "4"
	p2.TextStyle = ui.NewStyle(ui.ColorRed)
	p2.SetRect(32, 0, 48, 3)
	p2.Border = true

	p3 := widgets.NewParagraph()
	p3.Title = "总专注时间"
	p3.Text = "458.8"
	p3.TextStyle = ui.NewStyle(ui.ColorRed)
	p3.SetRect(0, 3, 15, 6)
	p3.Border = true

	p4 := widgets.NewParagraph()
	p4.Title = "本周专注时间"
	p4.Text = "4.5"
	p4.SetRect(15, 3, 31, 6)
	p4.TextStyle = ui.NewStyle(ui.ColorRed)
	p4.Border = true

	p5 := widgets.NewParagraph()
	p5.Title = "今日专注时间"
	p5.Text = "2"
	p5.TextStyle = ui.NewStyle(ui.ColorRed)
	p5.SetRect(32, 3, 48, 6)
	p5.Border = true

	p6 := widgets.NewParagraph()
	p6.Title = "总完成任务"
	p6.Text = "39"
	p6.TextStyle = ui.NewStyle(ui.ColorRed)
	p6.SetRect(0, 6, 15, 9)
	p6.Border = true

	p7 := widgets.NewParagraph()
	p7.Title = "本周完成任务"
	p7.Text = "0"
	p7.TextStyle = ui.NewStyle(ui.ColorRed)
	p7.SetRect(15, 6, 31, 9)
	p7.Border = true

	p8 := widgets.NewParagraph()
	p8.Title = "今日完成任务"
	p8.Text = "0"
	p8.TextStyle = ui.NewStyle(ui.ColorRed)
	p8.SetRect(32, 6, 48, 9)
	p8.Border = true

	ui.Render(p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func drawLine() {

}

func Draw() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	go drawText()
	// go drawLine()

	uiEvents := ui.PollEvents()
	for {

		switch ev := <-uiEvents; ev.ID {
		case "q", "<C-c>":
			return // break is not working
		}

	}

}
