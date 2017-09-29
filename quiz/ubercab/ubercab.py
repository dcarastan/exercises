#! /usr/bin/env python

# Uber's original name was "UberCab". Let's say we have a bunch of bumper stickers left over from those days
# which say "UBERCAB" and we decide to cut these up into their separate letters to make new words.
# So, for example, one sticker would give us the letters "U", "B", "E", "R", "C", "A", "B",
# which we could rearrange into other word[s] (like "car", "bear", etc)

# Challenge:
# Write a function that takes as its input an arbitrary string and as output,
# returns the number of intact "UBERCAB" stickers we would need to cut up to recreate that string.

# For instance:

# ubercab_stickers("car brace") would return "2", since we would need to cut up 2 stickers to provide enough letters to write "car brace"

import sys

UBERCAB = "UberCab"

def ubercab_stickers(phrase):
    letter_usage = {}
    sticker_count = 0
    for letter in phrase:
        if letter == " ":
            continue
        letter = letter.lower()
        if letter in UBERCAB.lower():
            if letter in letter_usage:
                letter_usage[letter] += 1

            stickers = (letter_usage[letter] / UBERCAB.lower().count(letter))
            if letter_usage[letter] % UBERCAB.lower().count(letter) > 0:
                stickers += 1
            if sticker_count < stickers:
                sticker_count = stickers
        else:
            return -1

    return sticker_count

def ubercab_stickers2(phrase):
    words = phrase.split(' ')
    letter_usage = {}
    for word in words:
        for letter in word.lower():
            if letter in UBERCAB.lower():
                if letter in letter_usage:
                    letter_usage[letter] += 1
            else:
                return -1

    # letter_usage has the letter usage distribution.
    sticker_usage = {}
    sticker_count = 0
    for letter in letter_usage:
        sticker_usage[letter] = letter_usage[letter] / UBERCAB.lower().count(letter)
        if sticker_count < sticker_usage[letter]:
            sticker_count = sticker_usage[letter]

    return sticker_count

#-------------------------------------------------------------------------------

if len(sys.argv) > 1:
    phrase = " ".join(sys.argv[1:])
else:
    phrase = "bbb"

print "computing '{}' sticker usage for: {}".format(UBERCAB,phrase)
stickers = ubercab_stickers(phrase)
if stickers > 0:
    print "it takes {} sticker(s) to print '{}'".format(stickers, phrase)
else:
    print "it can't be done. '{}' has extra letters not found in '{}'".format(phrase, UBERCAB)
