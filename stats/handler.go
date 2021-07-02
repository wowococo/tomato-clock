package stats

import (
	"context"
	_ "log"
	"termdash/container/grid"
	"termdash/widgets/button"
	_ "tomato-clock/sqliteopt"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
)

type widgets struct {
	lc     *linechart.LineChart
	t      *staticText
	button *layoutButtons
}

func inputs() []float64 {
	var values []float64
	//todo
	for i := 0; i < 200; i++ {
		values = append(values, float64(i))
	}
	return values
}

func newLineChart() (*linechart.LineChart, error) {
	lc, err := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorDefault)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorBlue)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorFuchsia)),
	)
	hdlerr(err)

	values := inputs()
	labels := map[int]string{}
	err = lc.Series("weektomato", values, linechart.SeriesXLabels(labels))
	return lc, err

}

type staticText struct {
	alltomatoT   *text.Text
	weektomatoT  *text.Text
	todaytomatoT *text.Text
	allftT       *text.Text
	weekftT      *text.Text
	todayftT     *text.Text
	alltaskT     *text.Text
	weektaskT    *text.Text
	todaytaskT   *text.Text
}

func newText() (*staticText, error) {
	alltmtT, err := text.New()
	err = alltmtT.Write("10.4", text.WriteCellOpts(cell.FgColor(cell.ColorDefault), cell.Bold()))
	if err != nil {
		return nil, err
	}

	wcT, err := text.New()
	err = wcT.Write("3", text.WriteCellOpts(cell.FgColor(cell.ColorDefault)))
	if err != nil {
		return nil, err
	}

	tcT, err := text.New()
	err = tcT.Write("3", text.WriteCellOpts(cell.FgColor(cell.ColorDefault)))
	if err != nil {
		return nil, err
	}

	allftT, err := text.New()
	err = allftT.Write("3", text.WriteCellOpts(cell.FgColor(cell.ColorDefault)))
	if err != nil {
		return nil, err
	}

	wftT, err := text.New()
	err = wftT.Write("3", text.WriteCellOpts(cell.FgColor(cell.ColorDefault)))
	if err != nil {
		return nil, err
	}

	tftT, err := text.New()
	err = tftT.Write("3", text.WriteCellOpts(cell.FgColor(cell.ColorDefault)))
	if err != nil {
		return nil, err
	}

	alltaskT, err := text.New()
	err = alltaskT.Write("3", text.WriteCellOpts(cell.FgColor(cell.ColorDefault)))
	if err != nil {
		return nil, err
	}

	wtaskT, err := text.New()
	err = wtaskT.Write("3", text.WriteCellOpts(cell.FgColor(cell.ColorDefault)))
	if err != nil {
		return nil, err
	}

	ttaskT, err := text.New()
	err = ttaskT.Write("3", text.WriteCellOpts(cell.FgColor(cell.ColorDefault)))
	if err != nil {
		return nil, err
	}

	return &staticText{
		alltomatoT:   alltmtT,
		weektomatoT:  wcT,
		todaytomatoT: tcT,
		allftT:       allftT,
		weekftT:      wftT,
		todayftT:     tftT,
		alltaskT:     alltaskT,
		weektaskT:    wtaskT,
		todaytaskT:   ttaskT,
	}, nil
}

type layoutButtons struct {
	dtomato *button.Button
	wtomato *button.Button
	mtomato *button.Button
	dtask   *button.Button
	wtask   *button.Button
	mtask   *button.Button
}

func newLayoutButtons(w widgets) {
	dtomato, err := button.New("每日番茄曲线", func() {
		return setLayout(w)
	})
	wtomato, err := button.New("每周番茄曲线", func() {
		return setLayout(w)
	})

}

type layoutType int

const (
	layoutdtomato layoutType = iota
	layoutwtomato
	layoutmtomato
	layoutdtask
	layoutwtask
	layoutmtask
)

func setLayout(w widgets, lt layoutType) {
	upcols := []grid.Element{
		grid.RowHeightPerc(20,
			grid.ColWidthPerc(11,
				grid.Widget(w.t.alltomatoT,
					container.BorderTitle("总完成番茄数"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light),
					container.AlignHorizontal(align.HorizontalRight))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.weektomatoT,
					container.BorderTitle("本周完成番茄数"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.todaytomatoT,
					container.BorderTitle("今日完成番茄数"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.allftT,
					container.BorderTitle("总专注时间"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.weekftT,
					container.BorderTitle("本周专注时间"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.todayftT,
					container.BorderTitle("今日专注时间"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.alltaskT,
					container.BorderTitle("总完成任务"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.weektaskT,
					container.BorderTitle("本周完成任务"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.todaytaskT,
					container.BorderTitle("今日完成任务"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
		),
		grid.RowHeightPerc(5,
			grid.ColWidthPerc(33, grid.Widget(w.button.dtomato)),
			grid.ColWidthPerc(33, grid.Widget(w.button.wtomato)),
			grid.ColWidthPerc(33, grid.Widget(w.button.mtomato)),
		),
		grid.RowHeightPerc(5,
			grid.ColWidthPerc(33, grid.Widget(w.button.dtask)),
			grid.ColWidthPerc(33, grid.Widget(w.button.wtask)),
			grid.ColWidthPerc(33, grid.Widget(w.button.mtask)),
		),
	}
	switch lt {
	case layoutdtomato:
		upcols = append(upcols, w.lc.dtmt)
	case layoutwtomato:
		upcols = append(upcols, w.lc.wtmt)
	case layoutmtomato:
		upcols = append(upcols, w.lc.mtmt)
	case layoutdtask:
		upcols = append(upcols, w.lc.dtask)
	case layoutwtask:
		upcols = append(upcols, w.lc.wtask)
	case layoutmtask:
		upcols = append(upcols, w.lc.mtask)

	}
	builder := grid.New()
	builder.Add(upcols...)
	opts, err := builder.Build()
	return
}

func newWidgets() *widgets {
	lc, err := newLineChart()
	t, err := newText()
	lb, err := newLayoutButtons()
	hdlerr(err)
	return &widgets{
		lc: lc,
		t:  t,
	}
}

const rootID = "root"

func Draw() {
	t, err := tcell.New()
	hdlerr(err)
	defer t.Close()

	c, err := container.New(t, container.ID(rootID))
	hdlerr(err)

	ctx, cancel := context.WithCancel(context.Background())

	w := newWidgets()

	builder := grid.New()
	builder.Add(
		grid.RowHeightPerc(20,
			grid.ColWidthPerc(11,
				grid.Widget(w.t.alltomatoT,
					container.BorderTitle("总完成番茄数"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light),
					container.AlignHorizontal(align.HorizontalRight))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.weektomatoT,
					container.BorderTitle("本周完成番茄数"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.todaytomatoT,
					container.BorderTitle("今日完成番茄数"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.allftT,
					container.BorderTitle("总专注时间"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.weekftT,
					container.BorderTitle("本周专注时间"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.todayftT,
					container.BorderTitle("今日专注时间"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.alltaskT,
					container.BorderTitle("总完成任务"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.weektaskT,
					container.BorderTitle("本周完成任务"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.todaytaskT,
					container.BorderTitle("今日完成任务"),
					container.BorderColor(cell.ColorCyan),
					container.Border(linestyle.Light))),
		),

		grid.RowHeightPerc(70,
			grid.Widget(w.lc,
				container.BorderColor(cell.ColorCyan),
				container.BorderTitle("番茄曲线(周)"))))

	gridOpts, err := builder.Build()
	hdlerr(err)

	err = c.Update(rootID, gridOpts...)
	hdlerr(ctx.Err())

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	err = termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter))
	hdlerr(err)

}

func hdlerr(err error) {
	if err != nil {
		panic(err)
	}
}
