// wordfreq -- Web page word counter
//
// Author: Doru Carastan (doru@rocketmail.com)
//
// Usage:
//   wordfreq [-help|-h]
//   wordfreq -urls=<comma-seperated-one-or-more-urls>
//
//   wordfreq -urls https://www.aporeto.com/caesars-story-microbursting-cloud-security-services/,https://www.aporeto.com/trust-centric-security-authentication-and-authorization/,https://www.aporeto.com/accelerating-business-devops-and-microservices-part-ii-running-safer/

package main

import (
	"flag"
	"fmt"
	"github.com/jaytaylor/html2text"
	"net/http"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	optUrls string
)

func init() {
	flag.StringVar(&optUrls, "urls", "", "Comma separated one or more URLs")
}

func main() {
	flag.Parse()
	requiredFlagCheck([]string{"urls"})

	urls := strings.Split(optUrls, ",")

	fmt.Println("Performing text analysis on:")
	for _, url := range urls {
		fmt.Printf("\t%s\n", url)
	}
	fmt.Println()

	// Channels
	chFinished := make(chan bool)
	defer close(chFinished)

	// Start the page worker goroutines
	for i := 0; i < len(urls); {
		if len(urls[i]) > 0 {
			go pageWorker(i, urls[i], chFinished)
			i++
		}
	}

	// Subscribe and process data
	c := 0
	for c < len(urls) {
		select {
		case <-chFinished:
			c++
		}
	}

	fmt.Println("\nDone!")
}

// Verifies that required flags are specified.
// TODO: See if there is a way to organically extend the flag package for it.
func requiredFlagCheck(required []string) {
	have := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { have[f.Name] = true })
	for _, arg := range required {
		if !have[arg] {
			// TODO: use log.Fatalf() ?
			fmt.Fprintf(os.Stderr, "ERROR: Missing required -%s argument\n",
				arg)
			flag.Usage()
			// Report same error code as flag.Parse()
			os.Exit(2)
		}
	}
}

// Split text into word array.
// https://github.com/manjuprasadsn/goprogramming/blob/master/word_frequency.go
func splitOnNonLetters(text string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(text, notALetter)
}

// Extract text from a given webpage and perform word frequency analysis on it.
// Input: wid        - worker id
//        url        - url to process
//        chFinished - channel to report when the job is done
func pageWorker(wid int, url string, chFinished chan bool) {
	fmt.Printf("[%d] Processing: %s\n", wid, url)

	defer func() {
		// Notify that we're done on this function exit
		chFinished <- true
	}()

	// Read the web page
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[%d] ERROR: Failed to read '%s'", wid, url)
		return
	}

	body := resp.Body
	defer body.Close()

	// Extract the text content from HTML page
	text, err := html2text.FromReader(body)
	if err != nil {
		panic(err)
	}

	// Calculate word frequency.
	frequencyForWord := make(map[string]int)
	for _, word := range splitOnNonLetters(text) {
		/* UTFMax = maximum number of bytes of a UTF-8 encoded Unicode char.
		   RuneCountInString returns the number of runes in p. Erroneous and
		   short encodings are treated as single runes of width 1 byte. */
		if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
			frequencyForWord[strings.ToLower(word)] += 1
		}
	}

	// Write url<N>.txt report file
	reportFile := fmt.Sprintf("url%d.txt", wid)

	f, err := os.Create(reportFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("url: " + url + "\n")
	if err != nil {
		panic(err)
	}

	for key, value := range frequencyForWord {
		line := fmt.Sprintf("  %s: %d\n", key, value)
		_, err := f.WriteString(line)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("[%d] Wrote: %s\n", wid, reportFile)
}
