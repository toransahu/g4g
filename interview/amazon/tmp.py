#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Fri Nov 17 09:10:18 2017

@author: toran
"""

m2 = [[1,2],[3,4]]

for i in range(2):
    for j in range(1,2):
        tmp = m2[i][j]
        m2[i][j] = m2[j][i]
        m2[j][i] = tmp

print(m2)
