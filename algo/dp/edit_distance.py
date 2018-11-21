#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2018-11-21 12:47

"""
Edit Distance.

USES
----
>> from edit_distance import edit_distance
>>
>> string_a = "abcd"
>> string_b = "abde"
>>
>> distance = edit_distance(string_a, string_b)
"""

from pprint import pprint

__author__ = "Toran Sahu <toran.sahu@yahoo.com>"
__license__ = "Distributed under terms of the MIT license"


class Algorithm:
    """Edit Distance Algorithm Class."""

    def __init__(self, str1, str2):
        """Initialize the Class.

        :param str1: String A
        :param str2: String B
        """
        self._str1 = str1
        self._str2 = str2
        self._len1 = len(self._str1)
        self._len2 = len(self._str2)
        self._distance_table = [
            [0 for _ in range(self._len1 + 1)] for _ in range(self._len2 + 1)
        ]
        for i in range(self._len1 + 1):
            self._distance_table[0][i] = i
        for j in range(self._len2 + 1):
            self._distance_table[j][0] = j

        pprint(self._distance_table)

    def _compute(self):
        """Compute the output."""
        if self._len1 == 0:
            return self._len2
        elif self._len2 == 0:
            return self._len1

        for i in range(1, self._len2 + 1):
            for j in range(1, self._len1 + 1):
                if self._str1[j - 1] == self._str2[i - 1]:
                    self._distance_table[i][j] = self._distance_table[i - 1][
                        j - 1
                    ]
                else:
                    self._distance_table[i][j] = 1 + min(
                        self._distance_table[i - 1][j - 1],
                        self._distance_table[i - 1][j],
                        self._distance_table[i][j - 1],
                    )
                pprint(self._distance_table)

        return self._distance_table[self._len2][self._len1]


def edit_distance(string_a, string_b):
    """Compute Edit Distance.

    :param string_a: String A
    :param string_b: String B
    """
    return Algorithm(string_a, string_b)._compute()
