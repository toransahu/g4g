"""Binary Tree."""

from queue import Queue


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
    Type: Depth First Traversal
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
    Type: Depth First Traversal
    """
    if root:
        result.append(root.data)
        pre_order(root.left, result)
        pre_order(root.right, result)
    return result


def post_order(root, result):
    """
    Post Order Traversal.

    Type: Depth First Traversal
    Prints binary tree.
    Order: Left, right, root
    Implementation: Recursive
    """
    if root:
        post_order(root.left, result)
        post_order(root.right, result)
        result.append(root.data)
    return result


def level_order_traversal(root):
    """
    Traverse in Breadth First Search order.
    * Desc:
            1. Find out total level of the tree : Traverse each sub-tree + Compare level of left & right
            2. Loop from first level to last
            3. For each level
    * Type: Breadth First Traversal
    * Approach: Iterative

    :param root: Root of the tree
    :param result: Result in list format
    :return: result
    """
    # find level of the tree
    l = level_of_tree(root)

    # for exclusive upper bound use l + 1
    for level in range(1, l+1):
        print_given_level(root, level)


def level_of_tree(root):
    """
    Find out total levels in a binary Tree.

    Implementation: Recursive

    :param root: root of the tree
    :return: total levels in given tree
    """
    if root:
        l = level_of_tree(root.left)
        r = level_of_tree(root.right)
        if l > r:
            return l + 1
        else:
            return r + 1
    else:
        return 0


def print_given_level(root, level):
    """
    Print nodes of a binary tree of a given level.

    * Desc:
        1. If root is empty, return; If level reached 1, print the data
        2. Make Recursive  call for each child by decreasing level by 1
    :param root:
    :param level:
    :return:
    """
    if root is None:
        return
    elif level == 1:
        print(root.data)
    elif level > 1:
        print_given_level(root.left, level - 1)
        print_given_level(root.right, level - 1)


def level_order_traversal_queue(root, queue):
    """
    First visit & print node, then enqueue their child.
    Dequeue one by one and call the same function for that node.

    Type: Breadth First Traversal
    Implementation: Recursive
    Data Structure Used: Queue

    :param root: Root of the binary tree
    :param queue: An empty queue
    :return: Nothing, Just Prints
    """
    if root:
        print(root.data)
        if root.left:
            queue.put(root.left)
        if root.right:
            queue.put(root.right)
        level_order_traversal_queue(queue.get(), queue)


root = Node(1)
root.left = Node(2)
root.right = Node(3)
root.left.left = Node(4)
root.left.right = Node(5)
root.right.right = Node(7)
root.right.right.right = Node(20)
# print(in_order(root, []))
# print(pre_order(root, []))
# print(post_order(root, []))
# level_order_traversal(root)
# print(level_of_tree(root))
# q = Queue()
# level_order_traversal_queue(root, q)