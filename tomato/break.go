package tomato

import (
	"context"
	"log"
	"time"

	"github.com/mum4k/termdash/cell"

	"github.com/mum4k/termdash"
	// "github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
)

func breaktime(d time.Duration, tomatoID int64) {
	echo()
	tbinit()
	countdown(d, tomatoID, true)

}

const breakID = "break"

func echo() {
	t, err := tcell.New()
	if err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer t.Close()

	c, err := container.New(t, container.ID(breakID))
	if err != nil {
		panic(err)
	}

	_, cancel := context.WithCancel(context.Background())
	borderless, err := text.New()
	if err != nil {
		panic(err)
	}

	if err := borderless.Write("TIME TO BREAK ^_^ ",
		text.WriteCellOpts(
			cell.Bold(),
			cell.Blink(),
			cell.FgColor(cell.ColorAqua))); err != nil {
		panic(err)
	}

	cols := []grid.Element{
		grid.ColWidthPerc(40),
		grid.ColWidthPerc(30,
			grid.RowHeightPerc(40),
			grid.RowHeightPerc(20, grid.Widget(borderless)),
			grid.RowHeightPerc(40)),
		grid.ColWidthPerc(30),
	}

	builder := grid.New()
	builder.Add(cols...)
	gridOpts, err := builder.Build()
	if err != nil {
		panic(err)
	}

	err = c.Update(breakID, gridOpts...)
	if err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	controller, err := termdash.NewController(t, c, termdash.KeyboardSubscriber(quitter))
	if err != nil {
		panic(err)
	}
	defer controller.Close()

	time.Sleep(2 * time.Second)

}
