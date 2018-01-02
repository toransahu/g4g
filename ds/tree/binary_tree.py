"""Binary Tree."""


class Node:
    """Node of a Binary Tree."""
    def __init__(self, val):
        self.data = val
        self.left = None
        self.right = None


"""
Traverse & print the Binary Tree.

1. In Order
2. Post Order
3. Pre Order
"""


def in_order(root):
    """
    In Order Traversal.

    Prints binary tree.
    Order: Left, root, right
    Implementation: Recursive
    """
    if root:
        # call left sub-tree
        in_order(root.left)
        # print current root data
        print(root.data, end=' ')
        # call right sub-tree
        in_order(root.right)


root = Node(1)
root.left = Node(2)
root.right = Node(3)
root.left.left = Node(4)
root.left.right = Node(5)

in_order(root)
