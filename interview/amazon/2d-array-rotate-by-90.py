"""
Rotate by 90 degree.
Solutions:
    1. With extri space, replace rows into columns
    2. Inplace, First transpose matrix & then reverse columns
"""

m3 = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]

m4 = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [13, 14, 15, 16]]

## Solution 2

print(m4)

## transpose via 00,11,22,33....
for i in range(len(m4)):
    for j in range(i, len(m4[0])):
        m4[i][j], m4[j][i] = m4[j][i], m4[i][j]

## reverse column
for i in range(len(m4[0])):
    for j in range(len(m4) // 2):
        m4[i][len(m4) - 1 - j], m4[i][j] = m4[i][j], m4[i][len(m4) - 1 - j]

print(m4)
