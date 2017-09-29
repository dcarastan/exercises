#!/usr/bin/env python

# this one looks like your scripts with argv
def print_two(*args):
	arg1, arg2 = args
	print "arg1: %r, arg2: %r" % (arg1, arg2)

# *args is pointless, we can just do thia
def print_two_again(arg1, arg2):
	print "arg1: %r, arg2: %r" % (arg1, arg2)

# this just takes one argument
def print_one(arg1):
	print "arg1: %r" % arg1

def print_none():
	print "I got nothing."


print_two("Zed", "Cucu")
print_two_again("zed", "cucu")
print_one("ONE")
print_none()
