#! /usr/bin/env python

"""
uniquify -- removes duplicated lines from a file

uniquify [--help|-h]
uniquify --file=<filename> --output=<output-filename> [-verbose]

"""

# PyLint workaround for: Invalid constant name "log" (invalid-name)
# pylint: disable=C0103

import argparse
import logging
import sys

# Setup console logging facility
log = logging.getLogger()
handler = logging.StreamHandler()
formatter = logging.Formatter('%(asctime)s %(levelname)-8s %(message)s')
handler.setFormatter(formatter)
log.addHandler(handler)
log.setLevel(logging.DEBUG)


def remove_duplicates(args):
    '''
    Removes duplicated lines from file
    '''
    if args.verbose:
        log.setLevel(logging.INFO)
        log.info("Reading from %s", args.file)

    with open(args.output, 'w+') as out:
        if args.verbose:
            log.info("Output set to %s", args.output)

        cache = set()

        # Read line by line
        for line in open(args.file):
            # If we have not seen it, add it to the cache and write it into the
            # output file. This method requires enough memory to hold the
            # uniquified file in memory. Memory footprint can be further
            # reduced by retaining each line's hashed value instead.
            if line not in cache:
                out.write(line)
                cache.add(line)

#
#  M A I N
#
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Trivial file content uniquify'
                                                 ' tool')
    parser.add_argument('-verbose', help='Verbose mode', action='store_true')

    # http://stackoverflow.com/questions/24180527/
    #        argparse-required-arguments-listed-under-optional-arguments
    required = parser.add_argument_group('required arguments')
    required.add_argument('--file', metavar='<filename>',
                          help='Input text file', required=True)
    required.add_argument('--output', metavar='<filename>',
                          help='Output text file', required=True)

    parsed_args = parser.parse_args()

    sys.exit(remove_duplicates(parsed_args))
