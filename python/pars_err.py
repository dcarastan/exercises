#!/usr/bin/env python

import re

BLOCK_START = "\| Block start"
ERROR_FAIL = "    \| ERROR: Failure"

def parse_log(log_file):
	d = {}
	with open(log_file, "r") as f:
		blkno = 0
		errno = 0
		for line in f:
			print "%d %s" % (blkno, line.rstrip())
			if re.match(ERROR_FAIL, line):
				errno += 1
                d['B{:d}'.format(blkno)] = "E{:d}".format(errno)
                print "ERROR", line
	return d

print parse_log("pars_err.dat")
