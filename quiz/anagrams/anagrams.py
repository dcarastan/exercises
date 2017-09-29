#! /usr/bin/env python

# Anagrams have the same letter set in different positions. The character
# ordinal sum is the same for 'car' and 'rac'.

# Compute leteer sum.
def letterWeight(w):
    return sum([ord(c) for c in w])

def find_anagrams(l):
    weight_list = map(letterWeight,l)
    # Build the word weight list, iterate on it, pick words for which we have
    # more than one identical weight value.
    return [l[i] for i,num in enumerate(weight_list) if weight_list.count(num) > 1]


words = ["car", "race", "rac", "ecar", "me", "em", "zoo"]

print "Word list: {}".format(words)
print "Anagrams list: {}".format(find_anagrams(words))
