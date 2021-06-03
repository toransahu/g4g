#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2021-06-03 16:32

"""Sol 21."""


__author__ = 'Toran Sahu <toran.sahu@yahoo.com>'
__license__ = 'Distributed under terms of the MIT license'


def is_prime(N):
    if N < 2:
        return False

    is_divisible = False
    for i in range(2, N):
        if N % i == 0:
            is_divisible = True
            break
    return not is_divisible


def sol_21(N):
    if N < 1:
        return -1
    if N == 1:
        return 0
    if is_prime(N):
        return N

    min_result = float('INF')
    curr_result = float('INF')

    for i in range(2, N):
        if N % i != 0:
            continue

        curr_result = i + (N/i)
        min_result = min(curr_result, min_result)
    return min_result


assert sol_21(-1) == -1
assert sol_21(0) == -1
assert sol_21(1) == 0
assert sol_21(2) == 2
assert sol_21(3) == 3
assert sol_21(4) == 4
assert sol_21(5) == 5
assert sol_21(6) == 5
assert sol_21(7) == 7
assert sol_21(8) == 6
assert sol_21(9) == 6
assert sol_21(10) == 7
assert sol_21(11) == 11
assert sol_21(12) == 7
assert sol_21(13) == 13
assert sol_21(14) == 9
assert sol_21(15) == 8
