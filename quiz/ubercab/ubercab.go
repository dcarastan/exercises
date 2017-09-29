/*
Uber's original name was "UberCab". Let's say we have a bunch of bumper stickers
left over from those days which say "UBERCAB" and we decide to cut these up into
their separate letters to make new words. So, for example, one sticker would
give us the letters "U", "B", "E", "R", "C", "A", "B", which we could rearrange
into other word[s] (like "car", "bear", etc)

Challenge:
Write a function that takes as its input an arbitrary string and as output,
returns the number of intact "UBERCAB" stickers we would need to cut up to
recreate that string.

For instance:
ubercab_stickers("car brace") would return "2", since we would need to cut up 2
stickers to provide enough letters to write "car brace"
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

// STICKER defines the sticker string.
const STICKER = "UberCab"

// NullWriter type is a mock object base.
type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func uberStickers(phrase string) (int, error) {
	letterUsage := make(map[rune]int)
	stickerCount := 0
	phrase = strings.ToLower(phrase)
	sticker := strings.ToLower(STICKER)
	for _, letter := range phrase {
		if letter == ' ' {
			continue
		}
		if strings.Contains(sticker, string(letter)) {
			if letterUsage[letter] > 0 {
				letterUsage[letter]++
			} else {
				letterUsage[letter] = 1
			}

			// Update statistics counters.
			stickers := letterUsage[letter] / strings.Count(sticker, string(letter))
			if letterUsage[letter]%strings.Count(sticker, string(letter)) > 0 {
				stickers++
			}
			if stickerCount < stickers {
				stickerCount = stickers
			}
			log.Printf("stickers: %d  letterUsage[%s]: %v\n", stickers, string(letter), letterUsage[letter])

		} else {
			return 0, fmt.Errorf("Letter %s not found in sticker", string(letter))
		}
	}
	return stickerCount, nil
}

func main() {
	phrasePtr := flag.String("phrase", "cab", "Phrase to analyze")
	verbosePtr := flag.Bool("verbose", false, "Verbose execution")
	flag.Parse()
	if !*verbosePtr {
		log.SetOutput(new(NullWriter))
	}

	log.Printf("Computing '%s' sticker usage for '%s'\n", STICKER, *phrasePtr)
	if n, err := uberStickers(*phrasePtr); err != nil {
		log.Fatalln("Phrase can't be typed using sticker letters")
	} else {
		fmt.Printf("It takes %d '%s' sticker(s) to spell '%s'\n", n, STICKER, *phrasePtr)
	}
}
