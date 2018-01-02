from ds.tree.binary_tree import in_order, Node


def test_in_order_positive():
    """
    Testing inorder traversal.
    """
    root = Node(1)
    root.left = Node(2)
    root.right = Node(3)
    root.left.left = Node(4)
    root.left.right = Node(5)
    result = []
    assert in_order(root, result) == [4, 2, 5, 1, 3]


def test_in_order_negative():
    root = Node(1)
    root.left = Node(2)
    root.right = Node(3)
    root.left.left = Node(4)
    root.left.right = Node(5)
    result = []
    assert in_order(root, result) != []
