#!/usr/bin/env python

# input:  1 2 3 5 6 7 8 9 10 11 12 13 15 16 17 ..... (x4, 4x, 4xx, 345, 444)
# output: 1 2 3 4 5 6 7 8 9  10 11 12 13 14 15 O(lg(n))
import re, time

def print_lucky_number(n):

    pattern = r'4'
    i = 0
    for n in range(1, n+1):
        if re.search(pattern, str(n)):
            i += 1
    return n - i


def print_lucky_number2(n):

    if n < 4:
        return n
    if n >= 4 and n < 10:
       return n - 1
 
    # Calculate 10^(numDigits-1) 
    largest_base_num = 1
    while n / largest_base_num >= 10:
        largest_base_num *= 10

    # Extract the left most digit
    left_most_digit = n / largest_base_num
 
    if left_most_digit != 4:
      return (print_lucky_number2(left_most_digit) * print_lucky_number2(largest_base_num - 1) + 
             print_lucky_number2(left_most_digit) + print_lucky_number2(n % largest_base_num))
    else:
      return print_lucky_number2(4 * largest_base_num - 1)


start = time.time()
print "Lucky %d is actually %d" % (57, print_lucky_number2(57))
print '%d s' % (time.time() - start)
