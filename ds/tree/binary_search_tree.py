#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2018-11-19 18:00

"""Binary Search Tree.
DESCRIPTION: Binary Search Tree (BST) is essentially a binary tree with one important feature of relative ordering of nodes.

Relative Ordering in BST: In a BST all nodes to the:
        - left of a node have values less than the value of the node
        - right of a node have values greater than the values of the Node

"""


from .binary_tree import in_order


__author__ = "Toran Sahu <toran.sahu@yahoo.com>"
__license__ = "Distributed under terms of the MIT license"


class TreeNode:

    """TreeNode Class to add Node to BST"""

    def __init__(self):
        pass
