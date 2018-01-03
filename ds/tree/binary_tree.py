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


def pre_order(root, result):
    """
    Pre Order Traversal.

    Prints binary tree.
    Order: root, left, right
    Implementation: Recursive
    """
    if root:
        result.append(root.data)
        pre_order(root.left, result)
        pre_order(root.right, result)
    return result


def post_order(root, result):
    """
    Post Order Traversal.

    Prints binary tree.
    Order: Left, right, root
    Implementation: Recursive
    """
    if root:
        post_order(root.left, result)
        post_order(root.right, result)
        result.append(root.data)
    return result


root = Node(1)
root.left = Node(2)
root.right = Node(3)
root.left.left = Node(4)
root.left.right = Node(5)

print(in_order(root, []))
print(pre_order(root, []))
print(post_order(root, []))
