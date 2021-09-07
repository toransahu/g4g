#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2021-09-06 11:28

"""Sol 2.2"""


__author__ = 'Toran Sahu <toran.sahu@yahoo.com>'
__license__ = 'Distributed under terms of the MIT license'


def is_num(ch):
    return str.isdigit(ch)


def decode_sub_str(s):
    num_stack = []
    str_stack = []
    digit = ''  # tmp var to collect contiguous digit chars
    alpha = ''  # tmp var to collect contiguous a-z chars

    # iterate through all the characters in the given encoded string
    for ch in s:
        # if its a digit, then it could be of N figures/size
        if is_num(ch):
            # so collect each digit char
            digit += ch
            # also, in case of nested encoded string
            # if there is collected alphabet from preceeding pattern
            if alpha:
                # then push that to the string stack
                str_stack.append(alpha)
                # nullify the alpha temp var
                alpha = ''
        # if its a opening bracket
        elif ch == '[':
            # then there was a number prefixed to it, and collected
            # push that in num stack
            num_stack.append(int(digit))
            # nullify the digit temp var
            digit = ''
        # if its a closing bracket
        elif ch == ']':
            # then there could be some string prefixed to it
            if alpha:
                # push that into the string stack
                str_stack.append(alpha)
                # nullify the alpha temp var
                alpha = ''

            # on hitting a closing bracket, pop top num from the stack
            num = num_stack.pop()
            # and pop top/recent string 
            alpha = str_stack.pop()
            # form a decoded string using those
            decoded = ''.join(num * [alpha])

            # if we're in nested situation, there could be existing string in string_stack
            if len(str_stack) > 0:
                # then append this decoded string from nested/sub-problem to the existing/outer string/problem
                str_stack.append(str_stack.pop() + decoded)
            else:
                # we'll keep the final result in string stack, so push the decoded string
                str_stack.append(decoded)
            # nullify the alpha temp var
            alpha = ''
        # otherwise its a alphabet
        else:
            # collect the char in alpha tmp var until we hit a closing bracket or a num
            alpha += ch
    # as the final output is in the string stack, pop it
    out = [str_stack.pop()]
    return ''.join(out)


assert decode_sub_str('2[a]3[ca]0[b]') == 'aacacaca'
assert decode_sub_str('2[e2[f]]') == 'effeff'
assert decode_sub_str('2[a]3[ca]0[b]2[e2[f]]') == 'aacacacaeffeff'
assert decode_sub_str('2[e2[f2[g]]]') == 'efggfggefggfgg'
