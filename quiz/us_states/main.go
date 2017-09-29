// Aporeto problem #1: US states file generator

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// STATES defines the US states
const STATES string = `Alabama
Alaska
Arizona
Arkansas
California
Colorado
Connecticut
Delaware
Florida
Georgia
Hawaii
Idaho
Illinois
Indiana
Iowa
Kansas
Kentucky
Louisiana
Maine
Maryland
Massachusetts
Michigan
Minnesota
Mississippi
Missouri
Montana
Nebraska
Nevada
New Hampshire
New Jersey
New Mexico
New York
North Carolina
North Dakota
Ohio
Oklahoma
Oregon
Pennsylvania
Rhode Island
South Carolina
South Dakota
Tennessee
Texas
Utah
Vermont
Virginia
Washington
West Virginia
Wisconsin
Wyoming`

// NullWriter type is a mock object base.
type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func main() {
	filePtr := flag.String("create-file", "", "Output file")
	promptPtr := flag.Bool("no-prompt", false, "Disables prompting")
	verbosePtr := flag.Bool("verbose", false, "Verbose execution")
	flag.Parse()

	if !*verbosePtr {
		log.SetOutput(new(NullWriter))
	}

	if len(*filePtr) == 0 {
		log.Fatal("ERROR: Specify the output file using the '--create-file' option")
	}

	if _, err := os.Stat(*filePtr); !os.IsNotExist(err) {
		log.Println("File already exists")

		if !*promptPtr {
			reader := bufio.NewReader(os.Stdin)
			for cmd := "n"; strings.ToUpper(cmd) != "Y\n"; {
				fmt.Print("File exists. Overwrite [y/N] ? ")
				cmd, _ = reader.ReadString('\n')
			}
		}
		log.Println("Removing old file", *filePtr)
		err = os.Remove(*filePtr)
		if err != nil {
			log.Fatal("Failed to remove file", err)
		}
		log.Println("File removed")
	}

	data := []byte(STATES)
	log.Println("Writing file", *filePtr)
	err := ioutil.WriteFile(*filePtr, data, 0644)
	if err != nil {
		log.Fatal("Failed to write file", err)
	}
	log.Println("File created")
}
