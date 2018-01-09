"""
Remove different type of duplicate objects from a list.
"""


def integer_from_list(inp):
    """
    Removes duplicate integers from a list.

    Time Complexity: O(n)
    Auxilary Space: O(max)
    :param inp:
    :return:
    """
    MAX = max(inp)
    count = [0]*(max + 1)
    for e in inp:
        count[e] += 1

    out = []
    idx = 0
    for e in inp:
        if count[e] > 0:
            out[idx] = e
            idx += 1
            count[e] = -1

    return out


def dicts_from_list(inp):
    """
    Removes duplicate dict from a list.

    Time Complexity: O(n2)
    Auxilary Space: O(n)
    :param inp: duplicates
    :return: unique
    """
    print("Original Data:", inp)
    out = []
    for e in inp:
        # not in time compl = O(n)
        # it is same as using another for loop, better to use binary search for lookup in O(log n)
        if e not in out:
            out.append(e)
    return out


l_inp = [{'name': 'T', 'last': 'S', 'age': 0},
         {'name': 'A', 'last': 'S', 'age': 1},
         {'name': 'T', 'last': 'S', 'age': 0}]
print(dicts_from_list(l_inp))
