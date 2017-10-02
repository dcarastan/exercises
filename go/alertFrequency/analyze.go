package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// analyze examines the log file slice and builds up a Hourlyevents map.
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

				log.Printf("time: %s, severity: %s  product: %s  application: %s\n",
					time, severity, product, application)

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

// logScanner scans a folder and all its sub-folders for files matching a given
// pattern. Returns a slice with file paths.
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
