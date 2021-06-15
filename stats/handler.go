package stats

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
)

func drawText() {
	p0 := widgets.NewParagraph()
	p0.Text = "总完成番茄数"
	p0.SetRect(0, 0, 15, 6)
	p0.Border = true

	p1 := widgets.NewParagraph()
	p1.Text = "本周完成番茄数"
	p1.SetRect(15, 0, 31, 6)
	p1.Border = true

	p2 := widgets.NewParagraph()
	p2.Text = "今日完成番茄数"
	p2.SetRect(35, 0, 55, 6)
	p2.Border = true

	p3 := widgets.NewParagraph()
	p3.Text = "总专注时间"
	p3.SetRect(0, 6, 15, 12)
	p3.Border = true

	p4 := widgets.NewParagraph()
	p4.Text = "本周专注时间"
	p4.SetRect(15, 6, 35, 12)
	p4.Border = true

	p5 := widgets.NewParagraph()
	p5.Text = "今日专注时间"
	p5.SetRect(35, 6, 55, 12)
	p5.Border = true

	ui.Render(p0, p1, p2, p3, p4, p5)
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
