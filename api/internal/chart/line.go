package chart

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var end = time.Now()
var start = end.AddDate(0, -6, 0)

func generateLineData() []opts.LineData {
	items := make([]opts.LineData, 0)
	for d := start; d.After(end) == false; d = d.Add(time.Hour * 12) {
		items = append(items, opts.LineData{Value: XY{d, rand.Intn(300)}})
	}
	return items
}

func (cc ChartConfig) NewTimeseriesLineChart() *charts.Line {
	bc := cc.BaseConfiguration
	xAxis := opts.XAxis{
		Name:         "Time",
		NameLocation: "middle",
		NameGap:      40,
		Type:         "time",
		AxisLabel: &opts.AxisLabel{
			Show: true,
		},
	}
	yAxis := opts.YAxis{
		Name:         "Measurements",
		NameLocation: "middle",
		NameGap:      40,
		Type:         "value",
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(bc.Initialization),
		charts.WithTitleOpts(bc.Title),
		charts.WithXAxisOpts(xAxis),
		charts.WithYAxisOpts(yAxis),
		charts.WithDataZoomOpts(bc.DataZoomList...),
		charts.WithToolboxOpts(bc.Toolbox),
		charts.WithTooltipOpts(bc.Tooltip),
	)

	if !bc.Animation {
		line.SetGlobalOptions(charts.WithAnimation())
	}

	return line
}

type LineExamples struct{}

func (LineExamples) Examples() {
	// The incoming playload will supply an array of plot configs and a time range
	// We should query all of the incoming plot configs and construct plots based on the result
	// We'll want to render these to jpeg as they come in so we don't eat up too much memory with all of the plot data
	// TODO need to figure out how much we can downsample
	// Then render to pdf and send back to client

	bc := charts.BaseConfiguration{
		Animation: false,
		Initialization: opts.Initialization{
			Width:  "2268px",
			Height: "1701px",
		},
		Legend: opts.Legend{
			Show:   true,
			Bottom: "0%",
			Y:      "center",
		},
		// DataZoomList: []opts.DataZoom{{Type: "slider"}},
		// Tooltip: opts.Tooltip{Show: true, Trigger: "axis"},
		// Toolbox: opts.Toolbox{Show: true},
	}
	cc := ChartConfig{
		BaseConfiguration: bc,
	}

	l := cc.NewTimeseriesLineChart()

	xyValues := make([][]opts.LineData, 3)
	for i := 0; i < 3; i++ {
		xyValues[i] = generateLineData()
	}
	for i, xy := range xyValues {
		l.AddSeries(fmt.Sprintf("timeseries %d", i+1), xy)
	}

	// line.SetSeriesOptions(
	// 	charts.WithLineChartOpts(opts.LineChart{
	// 		ShowSymbol: false,
	// 	}),
	// 	charts.WithLineStyleOpts(opts.LineStyle{
	// 		Color: "rgb(255, 70, 131)",
	// 	}),
	// )

	page := components.NewPage()

	page.AddCharts(l)

	// var buf bytes.Buffer
	// bw := bufio.NewWriter(&buf)

	f, err := os.Create("examples/html/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))

	// pdf := gopdf.GoPdf{}
	// pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	// pdf.AddPage()
	// pdf.SetXY(250, 200)
	//
	// pdf.WritePdf("image.pdf")
}

// type LineExamples struct{}

// func (LineExamples) Examples() {
// 	page := components.NewPage()
// 	page.AddCharts(
// 	)
// 	f, err := os.Create("examples/html/line.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	page.Render(io.MultiWriter(f))
// }
