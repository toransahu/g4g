"""
Testing Binary Tree

Approach: setup & teardown (classic xunit style), module level
Source: https://docs.pytest.org/en/latest/xunit_setup.html
"""
import sys, os

myPath = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, myPath + "/../")


try:
    from tree.binary_tree import in_order, Node
except:
    from ds.tree.binary_tree import in_order, Node

root = None
result = None


# setting up a binary tree
# its root, and an empty result list
def setup_module(module):
    global root, result
    root = Node(1)
    root.left = Node(2)
    root.right = Node(3)
    root.left.left = Node(4)
    root.left.right = Node(5)
    result = []


def teardown_module(module):
    # have nothing to clean, so just passing
    pass


def test_in_order_positive():
    """
    Testing inorder traversal.
    """
    assert in_order(root, result) == [4, 2, 5, 1, 3]


def test_in_order_negative():
    assert in_order(root, result) != []
