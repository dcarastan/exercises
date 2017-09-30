// Group log entries of *.log files by product and hour
// HH:MM:SS Message with spaces SEVERITY PRODUCT APPLICATION
// 09:05:26 Something happened EXCEPTION Foo Bar

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/wcharczuk/go-chart"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// NullWriter type is a mock object base.
type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

// ProductEvents counts events per product
type ProductEvents map[string]int

// HourlyEvents maps product events to hours
type HourlyEvents map[int]ProductEvents

func analyze(logs []string) (HourlyEvents, error) {
	he := make(HourlyEvents)
	for _, f := range logs {
		log.Println("Reading:", f)
		fin, err := os.Open(f)
		if err != nil {
			log.Print("ERROR: Failed to open file: " + f)
			return nil, err
		}
		defer fin.Close()
		scanner := bufio.NewScanner(fin)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			var hour int
			line := strings.Split(scanner.Text(), " ")
			// The message contains log entry delimiters. In a real world the format
			// would have to be changed. For the purpose of this exercise workaround
			// by indexing the severity field from the end of the record.
			severity := line[len(line)-3 : len(line)-2][0]

			if severity == "EXCEPTION" {
				product := line[len(line)-2 : len(line)-1][0]
				application := line[len(line)-1 : len(line)][0]

				time := line[0]
				timeRecord := strings.Split(time, ":")
				hour, err = strconv.Atoi(timeRecord[0])
				if err != nil {
					return nil, err
				}

				log.Printf("time: %s, severity: %s  product: %s  application: %s\n", time, severity, product, application)

				if he[hour] == nil {
					he[hour] = make(ProductEvents)
				}
				he[hour][product]++
			}
		}
		if err = scanner.Err(); err != nil {
			return nil, err
		}
	}
	return he, nil
}

func logScanner(dir, filePattern string) ([]string, error) {
	var fl []string
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Print(err) // can't walk here
			return nil     // keep going
		}
		if f.IsDir() {
			return nil
		}
		var matched bool
		matched, err = filepath.Match(filePattern, f.Name())
		if err != nil {
			log.Print(err) // malformed pattern
			return err     // this is fatal.
		}
		if matched {
			log.Print("Found: " + path)
			fl = append(fl, path)
		}
		return nil
	})
	return fl, err
}

func appendIfMissing(slice []string, line string) []string {
	for _, v := range slice {
		if v == line {
			return slice
		}
	}
	return append(slice, line)
}

func main() {
	verbosePtr := flag.Bool("verbose", false, "Verbose execution")
	flag.Parse()
	if !*verbosePtr {
		log.SetOutput(new(NullWriter))
	}
	var err error
	dir := flag.Arg(0)
	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
	}

	logs, err := logScanner(dir, "*.log")
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	stats, err := analyze(logs)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	fmt.Printf("\nStats: %v\n", stats)

	http.HandleFunc("/", drawChart(stats))
	fmt.Println("Graph available at http://localhost:8890  Type Ctrl-C to exit...")
	log.Fatal(http.ListenAndServe(":8890", nil))
}

func drawChart(data HourlyEvents) func(http.ResponseWriter, *http.Request) {
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

		fmt.Printf("%+v\n", bars)

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

		// graph.Elements = []chart.Renderable{chart.Legend(&graph)}

		w.Header().Set("Content-Type", "image/png")
		err := graph.Render(chart.PNG, w)
		if err != nil {
			log.Fatal("Failed to render chart: ", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
