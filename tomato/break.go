package tomato

import (
	"context"
	"image"
	"log"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
)

func breaktime(d time.Duration, tomatoID int64) {
	echo()
	tbinit()
	countdown(d, tomatoID, true)

}

func echo() {
	t, err := tcell.New()
	if err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer t.Close()

	_, cancel := context.WithCancel(context.Background())
	borderless, err := text.New()
	if err != nil {
		panic(err)
	}
	borderless.Options().MaximumSize = image.Point{X: 10, Y: 10}

	if err := borderless.Write("TIME TO BREAK ^_^ "); err != nil {
		panic(err)
	}

	c, err := container.New(t,
		container.PlaceWidget(borderless),
		container.AlignHorizontal(align.HorizontalCenter),
		container.AlignVertical(align.VerticalMiddle))
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
