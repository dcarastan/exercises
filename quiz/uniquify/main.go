package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

// NullWriter type is a mock object base.
type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func appendIfMissing(slice []string, line string) []string {
	for _, v := range slice {
		if v == line {
			return slice
		}
	}
	return append(slice, line)
}

func removeDuplicates(in, out string) error {
	log.Println("Reading from", in)
	fin, err := os.Open(in)
	if err != nil {
		log.Fatal("Failed to open input file: " + in)
	}
	defer fin.Close()
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)

	ts := []string{}
	for scanner.Scan() {
		ts = appendIfMissing(ts, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fout, err := os.Create(out)
	if err != nil {
		log.Fatal("Failed to create output file" + out)
	}
	defer fout.Close()

	for _, l := range ts {
		_, err := fout.Write([]byte(l + "\n"))
		if err != nil {
			log.Fatal("Failed to write to output file" + out)
		}
	}
	return nil
}

func wordFrequencyAnalisys(in string) (map[string]int, error) {
	log.Println("Performing word analysis of", in)
	fin, err := os.Open(in)
	if err != nil {
		log.Fatal("ERROR: Failed to open file")
	}
	defer fin.Close()
	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanWords)
	words := make(map[string]int)
	for scanner.Scan() {
		w := scanner.Text()
		words[w]++
	}
	return words, nil
}

func printWordFrequency(words map[string]int) {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i int, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	for _, kv := range ss {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}

func main() {
	inPtr := flag.String("in", "", "Input file")
	outPtr := flag.String("out", "", "Output file")
	verbosePtr := flag.Bool("verbose", false, "Verbose execution")
	flag.Parse()
	if len(*inPtr) == 0 || len(*outPtr) == 0 {
		log.Fatalln("ERROR: Use --in and --out to specify the input and output files.")
		os.Exit(1)
	}
	if !*verbosePtr {
		log.SetOutput(new(NullWriter))
	}

	words, err := wordFrequencyAnalisys(*inPtr)
	if err != nil {
		log.Fatalln("ERROR: Failed to rank words")
		os.Exit(2)
	}
	printWordFrequency(words)

	if err := removeDuplicates(*inPtr, *outPtr); err != nil {
		log.Fatalln("ERROR: Failed to remove duplicates")
		os.Exit(3)
	}
}
