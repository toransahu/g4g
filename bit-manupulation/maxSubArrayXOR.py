#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Sat Nov 18 16:26:09 2017

@author: toran

Approach:
        1. O(n**2): Using two loops
        2. O(n): Using Trie (Pending)

url: http://www.geeksforgeeks.org/find-the-maximum-subarray-xor-in-a-given-array/        
"""

def get_max_sub_array(l):
    """Solution 1."""
    maxval = 0
    subarr = []
    for i in range(0, len(l)):
        maxtill = 0
        tmparr = []
        ## check from index 0 to n, next index 1 to n, next 2 to n and so on
        for j in range(i,len(l)):
            jv = l[j]
            maxtill = maxtill ^ jv
            tmparr.append(jv)
            ## maxval = max(maxtill, maxval)
            ## OR
            if maxtill > maxval:
                maxval = maxtill
                
                subarr = tmparr.copy()
    print(subarr)
    return maxval

l1 = [1, 2, 3, 4]        
l2 = [8, 1, 7, 12, 6, 2]
l3 = [4, 6]

print(get_max_sub_array(l2))        