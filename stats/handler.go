package stats

import (
	"context"
	_ "fmt"
	_ "time"
	"tomato-clock/sqliteopt"

	"github.com/mum4k/termdash"
	_ "github.com/mum4k/termdash/align"
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
	lc     *lCharts
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

// func dtmtInputs() []float64 {
// 	var (
// 		values []float64
// 		tmtLC  sqliteopt.TomatoLC
// 	)
// 	tmtLC.Query()
// }

type lCharts struct {
	dtmt  *linechart.LineChart
	wtmt  *linechart.LineChart
	mtmt  *linechart.LineChart
	dtask *linechart.LineChart
	wtask *linechart.LineChart
	mtask *linechart.LineChart
}

func newLineCharts() (*lCharts, error) {
	opts := []linechart.Option{
		linechart.AxesCellOpts(cell.FgColor(cell.ColorDefault)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorBlue)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorFuchsia)),
	}
	dtmtLC, err := linechart.New(opts...)
	if err != nil {
		return nil, err
	}

	values := inputs()
	labels := map[int]string{}
	err = dtmtLC.Series("daytomato", values, linechart.SeriesXLabels(labels))

	wtmtLC, err := linechart.New(opts...)
	if err != nil {
		return nil, err
	}
	err = wtmtLC.Series("weektomato", values, linechart.SeriesXLabels(labels))

	mtmtLC, err := linechart.New(opts...)
	if err != nil {
		return nil, err
	}
	err = mtmtLC.Series("monthtomato", values, linechart.SeriesXLabels(labels))

	dtaskLC, err := linechart.New(opts...)
	if err != nil {
		return nil, err
	}
	err = dtaskLC.Series("daytask", values, linechart.SeriesXLabels(labels))

	wtaskLC, err := linechart.New(opts...)
	if err != nil {
		return nil, err
	}
	err = wtaskLC.Series("weektask", values, linechart.SeriesXLabels(labels))

	mtaskLC, err := linechart.New(opts...)
	if err != nil {
		return nil, err
	}
	err = mtaskLC.Series("monthtask", values, linechart.SeriesXLabels(labels))

	return &lCharts{
		dtmt:  dtmtLC,
		wtmt:  wtmtLC,
		mtmt:  mtmtLC,
		dtask: dtaskLC,
		wtask: wtaskLC,
		mtask: mtaskLC,
	}, nil
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

const (
	tamatoTable = "tomato"
	taskTable   = "task"
)

const (
	tomatoColPgs = "progress"
	tomatoColTf  = "timefocused"
)

const (
	allTime    = "all"
	today      = "today"
	thisweek   = "thisweek"
	untilToday = "untiltoday"
	untilWeek  = "untilweek"
	untilMonth = "untilmonth"
)

func newText() (*staticText, error) {
	var mtc sqliteopt.Metric
	v0 := mtc.Query(tamatoTable, tomatoColPgs, allTime)
	alltmtT, err := text.New()
	err = alltmtT.Write(" "+v0, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v1 := mtc.Query(tamatoTable, tomatoColPgs, thisweek)
	wtmtT, err := text.New()
	err = wtmtT.Write(v1, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v2 := mtc.Query(tamatoTable, tomatoColPgs, today)
	ttmtT, err := text.New()
	err = ttmtT.Write(v2, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v3 := mtc.Query(tamatoTable, tomatoColTf, allTime)
	allftT, err := text.New()
	err = allftT.Write(v3, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v4 := mtc.Query(tamatoTable, tomatoColTf, thisweek)
	wftT, err := text.New()
	err = wftT.Write(v4, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v5 := mtc.Query(tamatoTable, tomatoColTf, today)
	tftT, err := text.New()
	err = tftT.Write(v5, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v6 := mtc.Query(taskTable, "", allTime)
	alltaskT, err := text.New()
	err = alltaskT.Write(v6, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v7 := mtc.Query(taskTable, "", thisweek)
	wtaskT, err := text.New()
	err = wtaskT.Write(v7, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v8 := mtc.Query(taskTable, "", today)
	ttaskT, err := text.New()
	err = ttaskT.Write(v8, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	return &staticText{
		alltomatoT:   alltmtT,
		weektomatoT:  wtmtT,
		todaytomatoT: ttmtT,
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

func newLayoutButtons(c *container.Container, w *widgets) (*layoutButtons, error) {
	opts := []button.Option{
		button.Height(1),
		button.WidthFor("每日番茄曲线"),
	}
	dtmt, err := button.New("每日番茄曲线", func() error {
		return setLayout(c, w, layoutdtomato)
	}, opts...)
	if err != nil {
		return nil, err
	}

	wtmt, err := button.New("每周番茄曲线", func() error {
		return setLayout(c, w, layoutwtomato)
	}, opts...)
	if err != nil {
		return nil, err
	}

	mtmt, err := button.New("每月番茄曲线", func() error {
		return setLayout(c, w, layoutmtomato)
	}, opts...)
	if err != nil {
		return nil, err
	}

	dtask, err := button.New("每日任务曲线", func() error {
		return setLayout(c, w, layoutdtask)
	}, opts...)
	if err != nil {
		return nil, err
	}

	wtask, err := button.New("每周任务曲线", func() error {
		return setLayout(c, w, layoutwtask)
	}, opts...)
	if err != nil {
		return nil, err
	}

	mtask, err := button.New("每月任务曲线", func() error {
		return setLayout(c, w, layoutmtask)
	}, opts...)
	if err != nil {
		return nil, err
	}

	return &layoutButtons{
		dtomato: dtmt,
		wtomato: wtmt,
		mtomato: mtmt,
		dtask:   dtask,
		wtask:   wtask,
		mtask:   mtask,
	}, nil
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

func setLayout(c *container.Container, w *widgets, lt layoutType) error {
	gridOpts, err := gridLayout(w, lt)
	if err != nil {
		return err
	}

	return c.Update(rootID, gridOpts...)

}

func gridLayout(w *widgets, lt layoutType) ([]container.Option, error) {
	upcols := []grid.Element{
		grid.RowHeightPerc(10,
			grid.ColWidthPerc(11,
				grid.Widget(w.t.alltomatoT,
					container.BorderTitle("总完成番茄数"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light),
					container.PaddingLeftPercent(3),
					container.PaddingRightPercent(3))),
			grid.ColWidthPerc(12,
				grid.Widget(w.t.weektomatoT,
					container.BorderTitle(" 本周完成番茄数"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(12,
				grid.Widget(w.t.todaytomatoT,
					container.BorderTitle(" 今日完成番茄数"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(10,
				grid.Widget(w.t.allftT,
					container.BorderTitle(" 总专注时间"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.weekftT,
					container.BorderTitle(" 本周专注时间"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.todayftT,
					container.BorderTitle(" 今日专注时间"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.alltaskT,
					container.BorderTitle(" 总完成任务"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.weektaskT,
					container.BorderTitle(" 本周完成任务"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light))),
			grid.ColWidthPerc(11,
				grid.Widget(w.t.todaytaskT,
					container.BorderTitle(" 今日完成任务"),
					container.BorderTitleAlignCenter(),
					container.Border(linestyle.Light))),
		),
		grid.RowHeightPerc(10,
			grid.ColWidthPerc(33, grid.Widget(w.button.dtomato)),
			grid.ColWidthPerc(33, grid.Widget(w.button.wtomato)),
			grid.ColWidthPerc(33, grid.Widget(w.button.mtomato)),
		),
		grid.RowHeightPerc(10,
			grid.ColWidthPerc(33, grid.Widget(w.button.dtask)),
			grid.ColWidthPerc(33, grid.Widget(w.button.wtask)),
			grid.ColWidthPerc(33, grid.Widget(w.button.mtask)),
		),
	}
	switch lt {
	case layoutdtomato:
		upcols = append(upcols,
			grid.RowHeightPerc(60,
				grid.Widget(w.lc.dtmt)))
	case layoutwtomato:
		upcols = append(upcols,
			grid.RowHeightPerc(60,
				grid.Widget(w.lc.wtmt)))
	case layoutmtomato:
		upcols = append(upcols,
			grid.RowHeightPerc(60,
				grid.Widget(w.lc.mtmt)))
	case layoutdtask:
		upcols = append(upcols,
			grid.RowHeightPerc(60,
				grid.Widget(w.lc.dtask)))
	case layoutwtask:
		upcols = append(upcols,
			grid.RowHeightPerc(60,
				grid.Widget(w.lc.wtask)))
	case layoutmtask:
		upcols = append(upcols,
			grid.RowHeightPerc(60,
				grid.Widget(w.lc.mtmt)))

	}
	builder := grid.New()
	builder.Add(upcols...)
	gridOpts, err := builder.Build()
	if err != nil {
		return nil, err
	}
	return gridOpts, nil
}

func newWidgets(c *container.Container) (*widgets, error) {
	lc, err := newLineCharts()
	if err != nil {
		return nil, err
	}

	t, err := newText()
	if err != nil {
		return nil, err
	}

	return &widgets{
		lc: lc,
		t:  t,
	}, nil
}

const rootID = "root"

func Draw() {
	t, err := tcell.New()
	hdlerr(err)
	defer t.Close()

	c, err := container.New(t, container.ID(rootID))
	hdlerr(err)

	ctx, cancel := context.WithCancel(context.Background())

	w, err := newWidgets(c)
	hdlerr(err)

	lb, err := newLayoutButtons(c, w)
	hdlerr(err)

	w.button = lb

	gridOpts, err := gridLayout(w, layoutdtomato)
	hdlerr(err)

	err = c.Update(rootID, gridOpts...)
	hdlerr(err)

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
