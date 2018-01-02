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


def in_order(root, result):
    """
    In Order Traversal.

    Prints binary tree.
    Order: Left, root, right
    Implementation: Recursive
    """
    if root:
        # call left sub-tree
        in_order(root.left, result)
        # store current root data
        result.append(root.data)
        # call right sub-tree
        in_order(root.right, result)
    return result


root = Node(1)
root.left = Node(2)
root.right = Node(3)
root.left.left = Node(4)
root.left.right = Node(5)
result = []
print(in_order(root, result))
