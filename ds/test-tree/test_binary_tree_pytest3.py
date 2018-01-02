"""
Testing Binary Tree

Framework: pytest
Approach:fixture with scope="class",  class based
"""

import pytest

try:
    from ..tree.binary_tree import in_order, Node
except:
    from ds.tree.binary_tree import in_order, Node


class TestBinaryTree:
    # setting up a binary tree
    # and returning its root
    @pytest.fixture(scope="class")
    def root(self):
        self.root = Node(1)
        self.root.left = Node(2)
        self.root.right = Node(3)
        self.root.left.left = Node(4)
        self.root.left.right = Node(5)
        return self.root

    # setting up an empty res list and returning it
    @pytest.fixture(scope="class")
    def result(self):
        """
        Module level fixture.

        Means it'll get created once, and will work for all the testcases.
        :return:
        """
        self.res = []
        return self.res

    def test_in_order_positive(self, root, result):
        assert in_order(root, result) == [4, 2, 5, 1, 3]

    def test_in_order_negative(self, root, result):
        assert in_order(root, result) != []
