#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Problem: Get binary representations of a number.

Approach: 
    1. Recursive
    2. Iterative
URL: https://www.hackerrank.com/challenges/coin-change

Created on Sat Nov 18 16:01:21 2017
@author: toran

"""


def get_binary_recursive(n):
    """ Recursive."""
    rem = []
    while n != 0:
        rem.append(n % 2)
        n = n // 2

    return "".join(map(str, rem[::-1]))


def get_binary_iterative(n):
    """
    Iterative.
    Check each ith bit if its 0(off) or 1(on).
    bitwise AND of 1&1 will give 1 and else will be 0.
    4 byte int i.e. 32 bits.
    """
    bi = [0] * 32

    for i in range(32):
        if ((2 ** i) & n) > 0:
            bi[i] = 1

    return "".join(map(str, bi[::-1]))


n = 10
print(get_binary_iterative(n))
print(get_binary_recursive(n))
