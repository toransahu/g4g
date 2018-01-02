import unittest
try:
    from .binary_tree import in_order, Node
except:
    from binary_tree import in_order, Node


class TestBinaryTree(unittest.TestCase):
    # def __init__(self):
    #     root = Node(1)
    #     root.left = Node(2)
    #     root.right = Node(3)
    #     root.left.left = Node(4)
    #     root.left.right = Node(5)
    #     self.root = root

    def test_in_order_traversal(self):
        self.assertEqual('4 2 5 1 3 ', '4 2 5 1 3 ', msg=None)


if __name__ == '__main__':
    unittest.main()
