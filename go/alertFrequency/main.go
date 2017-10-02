// Report log exception entries from *.log files by product and hour.
// Graph the results.
// File schema:
// HH:MM:SS Message with spaces SEVERITY PRODUCT APPLICATION
// 09:05:26 Something happened EXCEPTION ProdFoo AppBar

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// NullWriter type is a mock object base.
type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

// ProductEvents counts events per product
type ProductEvents map[string]int

// HourlyEvents maps product events to hours.
type HourlyEvents map[int]ProductEvents

func main() {
	serverURL := "http://localhost"
	serverPort := 8265

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

	// Extract product names.
	// Use a map for O(1) product de-dup. Empty structs take no data storage.
	products := make(map[string]struct{})
	for i := 0; i < 24; i++ {
		for k := range stats[i] {
			products[k] = struct{}{}
		}
	}

	for product := range products {
		// Setup product specific handlers
		http.HandleFunc("/stats/"+strings.ToLower(product), drawProductChart(stats, product))
		fmt.Printf("Access stats for product %s are available at %s:%d/stats/%s\n",
			product, serverURL, serverPort, strings.ToLower(product))
	}

	// TODO: Make / serve a templated index page referring to above product handles.
	http.HandleFunc("/", drawCombinedChart(stats))
	fmt.Printf("Access combined product stats are available at %s:%d\n", serverURL, serverPort)
	fmt.Println("Type Ctrl-C to exit...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
}
