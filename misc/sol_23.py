#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2021-06-07 19:05

"""Sol 23."""


__author__ = 'Toran Sahu <toran.sahu@yahoo.com>'
__license__ = 'Distributed under terms of the MIT license'


def sol_23(arr, n):
    if n <= 1:
        return arr

    l = 0
    r = n-1

    zero_ends_on = -1

    while(l < r):
        if arr[l] == 0:
            l += 1
            zero_ends_on = max(zero_ends_on, l)
            continue
        if arr[r] != 0:
            r -= 1
            continue
        if arr[l] != 0 and arr[r] == 0:
            tmp = arr[l]
            arr[l] = 0
            zero_ends_on = max(zero_ends_on, l)
            arr[r] = tmp
            l += 1
            r -= 1
    l = zero_ends_on + 1
    r = n-1
    while(l < r):
        if arr[l] == 1:
            l += 1
            continue
        if arr[r] == 2:
            r -= 1
            continue
        if arr[l] == 2 and arr[r] == 1:
            arr[l] = 1
            arr[r] = 2
            l += 1
            r -= 1
    return arr


assert sol_23([], 0) == []
assert sol_23([1, 0], 2) == [0, 1]
assert sol_23([1, 0, 1], 3) == [0, 1, 1]
assert sol_23([1, 0, 1, 0, 0, 1], 6) == [0, 0, 0, 1, 1, 1]

# as well as
assert sol_23([], 0) == []
assert sol_23([1, 0, 2], 3) == [0, 1, 2]
assert sol_23([2, 0, 1], 3) == [0, 1, 2]
assert sol_23([2, 0, 1, 2, 0, 1, 1], 7) == [0, 0, 1, 1, 1, 2, 2]
