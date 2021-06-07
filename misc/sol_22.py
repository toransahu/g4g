#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2021-06-07 18:45

"""Sol 22."""


__author__ = 'Toran Sahu <toran.sahu@yahoo.com>'
__license__ = 'Distributed under terms of the MIT license'


def sol_22(arr, n):
    if n <= 1:
        return arr

    l = 0
    r = n-1

    while(l < r):
        if arr[l] == 0:
            l += 1
        if arr[r] == 1:
            r -= 1
        if arr[l] == 1 and arr[r] == 0:
            arr[l] = 0
            arr[r] = 1
            l += 1
            r -= 1
    return arr


assert sol_22([], 0) == []
assert sol_22([1, 0], 2) == [0, 1]
assert sol_22([1, 0, 1], 3) == [0, 1, 1]
assert sol_22([1, 0, 1, 0, 0, 1], 6) == [0, 0, 0, 1, 1, 1]
