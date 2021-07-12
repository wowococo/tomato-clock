package stats

import (
	"context"
	"fmt"
	"strings"
	"time"
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

type layoutType int

const (
	layoutdtomato layoutType = iota
	layoutwtomato
	layoutmtomato
	layoutdtask
	layoutwtask
	layoutmtask
)

type widgets struct {
	lc     *lCharts
	t      *staticText
	button *layoutButtons
}

var tmtLC sqliteopt.TomatoLC

// daily tomato linechart
func dtmtInputs() ([]float64, map[int]string) {
	var (
		values  []float64
		dates   = make(map[string]int)
		XLabels = make(map[int]string)
	)
	now := time.Now()
	y, M, d, location := now.Year(), now.Month(), now.Day(), now.Location()
	start := time.Date(y, M-1, d, 0, 0, 0, 0, location)
	end := now
	diff := end.Sub(start)
	diffdays := int(diff.Hours() / 24)

	midays := diffdays / 2
	mid := time.Date(y, M-1, d+midays, 0, 0, 0, 0, location)

	st := start
	for i := 0; i <= diffdays; i++ {
		XLabels[i] = " "
		// date init in every loop?
		date := strings.Split(st.Format(time.RFC3339), "T")[0]
		values = append(values, 0)
		dates[date] = i
		st = st.AddDate(0, 0, 1)
	}

	XLabels[0] = fmt.Sprintf("%v %v", start.Month(), start.Day())
	XLabels[midays] = fmt.Sprintf("%v %v", mid.Month(), mid.Day())
	XLabels[diffdays] = "today"

	res := tmtLC.Query(tamatoTable, untilToday).(map[string]float64)
	for k, v := range res {
		if i, ok := dates[k]; ok {
			values[i] = v
		}
	}

	return values, XLabels
}

// weekly tomato linechart
func wtmtInputs() ([]float64, map[int]string) {
	var (
		values  []float64
		dates   = make(map[string]int)
		XLabels = make(map[int]string)
	)
	// mon of this week of this time
	mon := func(t time.Time) time.Time {
		weekday := t.Weekday()
		mondate := t.AddDate(0, 0, int(time.Monday-weekday))
		return mondate
	}

	now := time.Now()
	mondate := mon(now)
	y, M, d := mondate.Date()
	location := mondate.Location()
	start := mon(time.Date(y, M-6, d, 0, 0, 0, 0, location))
	// syear, sweek := start.ISOWeek()
	end := mondate
	// eyear, eweek := end.ISOWeek()
	diff := end.Sub(start)
	diffdays := int(diff.Hours() / 24)
	diffweeks := diffdays / 7

	midays := diffdays / 2
	midweeks := diffweeks / 2 
	mid := mon(start.AddDate(0, 0, midays))
	

	st := start
	for i := 0; i <= diffweeks; i++ {
		XLabels[i] = " "
		year, week := st.ISOWeek()
		w := string(week)
		var date string
		if len(w) == 1 {
			date = fmt.Sprintf("%d-0%s", year, w)
		}
		if len(w) == 2 {
			date = fmt.Sprintf("%d-%s", year, w)
		}

		values = append(values, 0)
		dates[date] = i
		st = st.AddDate(0, 0, 7)
	}

	endOfWeek := func(start time.Time) time.Time {
		return start.AddDate(0, 0, 6)
	}
	startSunday := endOfWeek(start)
	midSunday := endOfWeek(mid)
	XLabels[0] = fmt.Sprintf(
		"%v %v-%v %v",
		start.Month(),
		start.Day(),
		startSunday.Month(),
		startSunday.Day())
	XLabels[midweeks] = fmt.Sprintf(
		"%v %v-%v %v",
		mid.Month(),
		mid.Day(),
		midSunday.Month(),
		midSunday.Day())
	XLabels[diffweeks] = fmt.Sprintf("%v %v-today", end.Month(), end.Day())

	res := tmtLC.Query(tamatoTable, untilWeek).(map[string]float64)
	for k, v := range res {
		if i, ok := dates[k]; ok {
			values[i] = v
		}
	}
	fmt.Println(values)
	time.Sleep(5 * time.Second)
	return values, XLabels
}

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

	values0, XLabels0 := dtmtInputs()
	err = dtmtLC.Series("daytomato", values0, linechart.SeriesXLabels(XLabels0))
	if err != nil {
		return nil, err
	}

	wtmtLC, err := linechart.New(opts...)
	if err != nil {
		return nil, err
	}

	values1, XLabels1 := wtmtInputs()
	err = wtmtLC.Series("weektomato", values1, linechart.SeriesXLabels(XLabels1))

	mtmtLC, err := linechart.New(opts...)
	if err != nil {
		return nil, err
	}

	values := []float64{0, 0, 0, 0, 0}
	var labels map[int]string	
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
	err = alltmtT.Write("    "+v0, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v1 := mtc.Query(tamatoTable, tomatoColPgs, thisweek)
	wtmtT, err := text.New()
	err = wtmtT.Write("    "+v1, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v2 := mtc.Query(tamatoTable, tomatoColPgs, today)
	ttmtT, err := text.New()
	err = ttmtT.Write("    "+v2, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v3 := mtc.Query(tamatoTable, tomatoColTf, allTime)
	allftT, err := text.New()
	err = allftT.Write("    "+v3, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v4 := mtc.Query(tamatoTable, tomatoColTf, thisweek)
	wftT, err := text.New()
	err = wftT.Write("    "+v4, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v5 := mtc.Query(tamatoTable, tomatoColTf, today)
	tftT, err := text.New()
	err = tftT.Write("    "+v5, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v6 := mtc.Query(taskTable, "", allTime)
	alltaskT, err := text.New()
	err = alltaskT.Write("    "+v6, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v7 := mtc.Query(taskTable, "", thisweek)
	wtaskT, err := text.New()
	err = wtaskT.Write("    "+v7, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		return nil, err
	}

	v8 := mtc.Query(taskTable, "", today)
	ttaskT, err := text.New()
	err = ttaskT.Write("    "+v8, text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
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
