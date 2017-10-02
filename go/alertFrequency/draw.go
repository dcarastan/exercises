package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart"
	"log"
	"net/http"
	"strconv"
)

// drawCombinedChart accumulates stats for all products.
func drawCombinedChart(data HourlyEvents) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bars []chart.StackedBar
		// Build a chart.StackedBarChart struct.
		for i := 0; i < 24; i++ {
			var barValues []chart.Value
			for k, v := range data[i] {
				barValues = append(barValues, chart.Value{Value: float64(v), Label: k})
			}
			sb := chart.StackedBar{
				Name:   strconv.Itoa(i),
				Values: barValues,
				Width:  1,
			}
			bars = append(bars, sb)
		}

		graph := chart.StackedBarChart{
			Title:      "Exceptions by Product by Hour Bar Chart",
			TitleStyle: chart.StyleShow(),
			Background: chart.Style{
				Padding: chart.Box{
					Top:    80,
					Left:   10,
					Right:  10,
					Bottom: 10,
				},
			},
			Height:     512,
			Width:      1024,
			BarSpacing: 39,
			XAxis: chart.Style{
				Show: true,
			},
			YAxis: chart.Style{
				Show: true,
			},
			Bars: bars,
		}

		w.Header().Set("Content-Type", "image/png")
		err := graph.Render(chart.PNG, w)
		if err != nil {
			log.Fatal("Failed to render chart: ", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

// drawProductChart generates a product graph for handled route.
func drawProductChart(data HourlyEvents, product string) func(http.ResponseWriter, *http.Request) {
	// This can be optionally handled by passing the product as an URL parameter.
	// A closure leaves no door for being 'probed' by interested parties.
	return func(w http.ResponseWriter, r *http.Request) {
		var bars []chart.Value
		for i := len(bars); i < 24; i++ {
			bars = append(bars, chart.Value{Value: float64(data[i][product]), Label: fmt.Sprintf("%d:00", i)})
		}

		graph := chart.BarChart{
			Title:      product + " exceptions by hour",
			TitleStyle: chart.StyleShow(),
			Background: chart.Style{
				Padding: chart.Box{
					Top:    60,
					Left:   10,
					Right:  10,
					Bottom: 20,
				},
			},
			Height:     512,
			Width:      1024,
			BarSpacing: 30,
			XAxis: chart.Style{
				Show: true,
			},
			YAxis: chart.YAxis{
				Style: chart.Style{
					Show: true,
				},
			},
			Bars: bars,
		}

		w.Header().Set("Content-Type", "image/png")
		err := graph.Render(chart.PNG, w)
		if err != nil {
			log.Fatal("Failed to render chart: ", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
