#! /usr/bin/env python
"""
    anagram2.py -- Trivial anagram checker

    Code can be profiled after 'pip install line_profiler' and placing
    @profile decorator on function of interest using:

    kernprof -l -v anagrams2.py
"""
import argparse
import time

LONG_ANAGRAM_THRESHOLD = 2


@profile
def anagram_check(words):
    """ Anagram checker

        words - list of words to process
    """
    anagrams = [[words[0]]]

    for i in range(1, len(words)):
        # Check to see if we have the word in the anagrams list
        found = False
        word1 = list(words[i])

        for anagram in anagrams:
            word2 = list(anagram[0])

            if sorted(word1) == sorted(word2):
                anagram.append(words[i])
                found = True

        if not found:
            # Start a new list for it
            anagrams.append([words[i]])

    return anagrams


class Timer(object):
    """ Code timing class
    """
    def __init__(self, verbose=False):
        self.secs = 0
        self.msecs = 0
        self.start = 0
        self.end = 0
        self.verbose = verbose

    def __enter__(self):
        self.start = time.time()
        return self

    def __exit__(self, *args):
        self.end = time.time()
        self.secs = self.end - self.start
        self.msecs = self.secs * 1000  # millisecs
        if self.verbose:
            print 'elapsed time: %f ms' % self.msecs

    def checkpoint(self):
        """ Returns elapsed time
        """
        return time.time() - self.start


def main(args):
    """ Main function
    """

    if args.input:
        # Use supplied word list file
        with open(args.input, 'r') as word_file:
            words = word_file.read().splitlines()
    else:
        words = ["dog", "god", "gin", "python", "rum", "typhon", "aeolian", "aeolina", "aeonial"]

    print "Processing word list ..."
    with Timer(verbose=True) as tic_tok:
        anagrams = anagram_check(words)

    if not args.input:
        print "Words: {}".format(words[:7])
        print "Result: {}".format(anagrams[:7])

    long_anagrams = []
    longest_anagram = []
    for anagram in anagrams:
        if len(anagram) > len(longest_anagram):
            longest_anagram = anagram
        if len(anagram) > LONG_ANAGRAM_THRESHOLD:
            long_anagrams.append(anagram)
    print "Anagrams chains longer than {}: {}".format(LONG_ANAGRAM_THRESHOLD, long_anagrams)
    print "Longest anagram chain: {}".format(longest_anagram)


#
#   M A I N
#
if __name__ == "__main__":

    parser = argparse.ArgumentParser(description='Trivial anagram checker')
    parser.add_argument('--input', help='One word per line word file.')

    main(parser.parse_args())
