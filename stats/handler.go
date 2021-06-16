package stats

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

/*
  15     16     16
-----|------|------|
| 3
| 3
-----|------|------|

*/

func drawText() {
	pk0 := widgets.NewParagraph()
	pk0.Text = "总完成番茄数"
	pk0.SetRect(0, 0, 15, 3)
	pk0.Border = true

	pv0 := widgets.NewParagraph()
	pv0.Text = "927.7"
	pv0.SetRect(0, 3, 15, 6)
	pv0.Border = true

	pk1 := widgets.NewParagraph()
	pk1.Text = "本周完成番茄数"
	pk1.SetRect(15, 0, 31, 3)
	pk1.Border = true

	pv1 := widgets.NewParagraph()
	pv1.Text = "9"
	pv1.SetRect(15, 3, 31, 6)
	pv1.Border = true

	pk2 := widgets.NewParagraph()
	pk2.Text = "今日完成番茄数"
	pk2.SetRect(32, 0, 48, 3)
	pk2.Border = true

	pv2 := widgets.NewParagraph()
	pv2.Text = "4"
	pv2.SetRect(32, 3, 48, 6)
	pv2.Border = true

	pk3 := widgets.NewParagraph()
	pk3.Text = "总专注时间"
	pk3.SetRect(0, 6, 15, 9)
	pk3.Border = true

	pv3 := widgets.NewParagraph()
	pv3.Text = "458.8"
	pv3.SetRect(0, 9, 15, 12)
	pv3.Border = true

	pk4 := widgets.NewParagraph()
	pk4.Text = "本周专注时间"
	pk4.SetRect(15, 6, 31, 9)
	pk4.Border = true

	pv4 := widgets.NewParagraph()
	pv4.Text = "4.5"
	pv4.SetRect(15, 9, 31, 12)
	pv4.Border = true

	pk5 := widgets.NewParagraph()
	pk5.Text = "今日专注时间"
	pk5.SetRect(32, 6, 48, 9)
	pk5.Border = true

	pv5 := widgets.NewParagraph()
	pv5.Text = "2"
	pv5.SetRect(32, 9, 48, 12)
	pv5.Border = true

	pk6 := widgets.NewParagraph()
	pk6.Text = "总完成任务"
	pk6.SetRect(0, 12, 15, 15)
	pk6.Border = true

	pv6 := widgets.NewParagraph()
	pv6.Text = "39"
	pv6.SetRect(0, 15, 15, 18)
	pv6.Border = true

	pk7 := widgets.NewParagraph()
	pk7.Text = "本周完成任务"
	pk7.SetRect(15, 12, 31, 15)
	pk7.Border = true

	pv7 := widgets.NewParagraph()
	pv7.Text = "0"
	pv7.SetRect(15, 15, 31, 18)
	pv7.Border = true

	pk8 := widgets.NewParagraph()
	pk8.Text = "今日完成任务"
	pk8.SetRect(32, 12, 48, 15)
	pk8.Border = true

	pv8 := widgets.NewParagraph()
	pv8.Text = "0"
	pv8.SetRect(32, 15, 48, 18)
	pv8.Border = true

	ui.Render(pk0, pv0, pk1, pv1, pk2, pv2, pk3, pv3, pk4, pv4, pk5, pv5, pk6, pv6, pk7, pv7, pk8, pv8)
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
