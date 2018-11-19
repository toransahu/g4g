"""
Testing Binary Tree

Framework: pytest
Approach: fixture with scope="module", function level
"""

import pytest
import sys, os

myPath = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, myPath + "/../")


try:
    from tree.binary_tree import in_order, Node
except:
    from ds.tree.binary_tree import in_order, Node


# setting up a binary tree
# and returning its root
@pytest.fixture(scope="module")
def root():
    root = Node(1)
    root.left = Node(2)
    root.right = Node(3)
    root.left.left = Node(4)
    root.left.right = Node(5)
    return root


# setting up an empty res list and returning it
@pytest.fixture(scope="module")
def result():
    """
    Module level fixture.

    Means it'll get created once, and will work for all the testcases.
    :return:
    """
    res = []
    return res


def test_in_order_positive(root, result):
    """
    Testing inorder traversal.

    :param root: its above defined method name, it will return 'root' of tree as param
    :param result: its above defined method name result, it will return 'res' list as param
    :return: nothing, asserts only
    """
    assert in_order(root, result) == [4, 2, 5, 1, 3]


def test_in_order_negative(root, result):
    assert in_order(root, result) != []
