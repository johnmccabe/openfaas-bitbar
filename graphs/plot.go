package graphs

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/color"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/common/model"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/brewer"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Only show important part of metric name
var labelText = regexp.MustCompile("\\{(.*)\\}")

const format string = "png"

// Number of data points for the plot
const step = 100

// Plot creates a plot from metric data and saves it to a temporary file.
// It's the callers responsibility to remove the returned file when no longer needed.
func Plot(metrics model.Matrix, title string) (io.WriterTo, error) {
	p, err := plot.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create new plot: %v", err)
	}

	// titleFont, err := vg.MakeFont("Helvetica-Bold", 4*vg.Millimeter)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to load font: %v", err)
	// }
	textFont, err := vg.MakeFont("Helvetica", 3*vg.Millimeter)
	if err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}

	// p.Title.Text = title
	// p.Title.Font = titleFont
	p.BackgroundColor = color.Transparent
	// p.Title.Padding = 2 * vg.Millimeter
	p.X.Tick.Marker = plot.TimeTicks{Format: "15:04"}
	p.X.Tick.Label.Font = textFont
	p.Y.Tick.Label.Font = textFont
	p.Y.Min = 0
	// p.Legend.Font = textFont
	// p.Legend.Top = true
	// p.Legend.YOffs = 15 * vg.Millimeter

	// Color palette for drawing lines
	paletteSize := 8
	palette, err := brewer.GetPalette(brewer.TypeAny, "Dark2", paletteSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get color palette: %v", err)
	}
	colors := palette.Colors()

	for s, sample := range metrics {
		data := make(plotter.XYs, len(sample.Values))
		for i, v := range sample.Values {
			data[i].X = float64(v.Timestamp.Unix())
			f, err := strconv.ParseFloat(v.Value.String(), 64)
			if err != nil {
				return nil, fmt.Errorf("sample value not float: %s", v.Value.String())
			}
			data[i].Y = f
		}

		l, err := plotter.NewLine(data)
		if err != nil {
			return nil, fmt.Errorf("failed to create line: %v", err)
		}
		l.LineStyle.Width = vg.Points(1)
		l.LineStyle.Color = colors[s%paletteSize]

		p.Add(l)
		if len(metrics) > 1 {
			m := labelText.FindStringSubmatch(sample.Metric.String())
			if m != nil {
				p.Legend.Add(m[1], l)
			}
		}
	}

	// Draw plot in canvas with margin
	margin := 2 * vg.Millimeter
	width := 6 * vg.Centimeter
	height := 4 * vg.Centimeter
	c, err := draw.NewFormattedCanvas(width, height, format)
	if err != nil {
		return nil, fmt.Errorf("failed to create canvas: %v", err)
	}
	p.Draw(draw.Crop(draw.New(c), margin, -margin, margin, -margin))

	return c, nil
}

// FunctionGraph
func FunctionGraph(baseUrl, functionName, q string) (string, error) {
	promURL := strings.TrimRight(baseUrl, "/")
	query := fmt.Sprintf(q, functionName)
	queryTime := time.Now()
	queryRange, _ := time.ParseDuration("15m")
	metrics, err := Metrics(promURL, query, queryTime, queryRange, step)
	if err != nil {
		return "", err
	}

	title := "Functions per second"
	plot, err := Plot(metrics, title)
	if err != nil {
		return "", err
	}

	plotBuf := bytes.NewBuffer(nil)
	plot.WriteTo(plotBuf)

	encodedPlot := base64.StdEncoding.EncodeToString(plotBuf.Bytes())
	return encodedPlot, nil
}
