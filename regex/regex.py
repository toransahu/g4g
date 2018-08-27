#! /usr/bin/env python3
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2018-08-26 18:08


"""
regex.py

This module covers all the fundamentals of regular expressions explained with examples and applications.

- Basic Patterns:
    - use of .
        - matches any character
    - use of *
        - matches 0 or more repetition of a char/set
    - use of +
        - matches 1 or more repetition of a char/set
    - use of ?
        - matches 0 or 1 repetition of a char/set
        - also works as a Non-greedy pattern (with repeaters)
            - means in string '<b>Hello</b>' pattern r'<.*>' will match the whole string instead of '<b>'
            - but if pattern is used as r'<.*?>' then will match '<b>' only
            - AKA pcre (Perl Compatible Regular Expression)
    - use of \
        - sign of speciality
        - used before any special chars, to match that char
    - use of \t, \r, \n
        - \t matches a tab
        - \r matches a carriage return (line break) in Mac, \n\r in Windows
        - \n matches a line break ( carriage return) in Linux & Windows
        - note: On "old" printers, \r sent the print head back to the start of the line, and \n advanced the paper by one line. Both were therefore necessary to start printing on the next line.
    - use of \b (start and end of word anchors)
        - matches the boundary between word and non-word chars
        - matches the position called word boundaries
            - match has zero length
            - usually before (including start of the line) and after a word
            -e.g. <here>apple<here>
    - use of \B
        - opposite of \b
        - matches every position where \b does not
    - use of \s
    - use of \d
    - use of \D
    - use of \w
    - use of \W
    - use of []
        - dash - inside []
        - dot . inside []
    - use of ()
    - use of ^ (Start of String Anchors)
        - with square bracket (set of chars)
    - use of $ (End of String Anchors)
    - use of |
    - use of syntax (? ... )
        - Lookarounds
            - +ve Lookahead
            - -ve Lookahead
            - +ve Lookbehind
            - -ve Lookbehind
        - Non-Capturing Group

- Builtin Functions:
    - re.search
    - re.group
    - re.findall
        - with files
        - with groups
    - re.sub
    - re.compile
    - extra options to above functions


- Examples:
    - Repetitions
    - Leftmost and largest
    - Square Brackets (Set of chars)
        - Email example
    - Group Extraction
    - Greedy vs Non-Greedy
    - Substitution

Credits:
    - https://www.rexegg.com/regex-disambiguation.html#lookarounds
    - www.regular-expressions.info
"""

__author__ = 'Toran Sahu  <toran.sahu@yahoo.com>'
__license__ = 'Distributed under terms of the AGPL license.'


import re

str = "This#is#$% a&%name%$ #"
r = re.sub(r'(?<=[\w])([\W]+)(?=[\w])', ' ', str)
print(r)
r = re.sub(r'(?<=[A-Za-z0-9])([^A-Za-z0-9]+)(?=[A-Za-z0-9])', ' ', str)
print(r)
r = re.sub(r'q(?=u)', ' ', 'quit')
print(r)


