#!/usr/bin/env python

def twoSum(a, s):
    for i in range(0, len(a)):
        for j in range(i+1, len(a)):
            if a[i] + a[j] == s:
                return True
    return False

# Sort the array first and iterate
def twoSum2(a, s):
    a.sort()

    l = 0
    r = len(a) - 1

    while l < r:
        if (a[l] + a[r] == s):
            return True
        elif (a[l] + a[r] < s):
            l += 1
        else:
            r -= 1

    return False


num_array = [9,4,2,1,1,3,2]
print("Input aray: {}".format(num_array))
for n in range(1, 30):
    print "%r %r %d" % (twoSum(num_array, n), twoSum2(num_array, n), n)
