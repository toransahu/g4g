import unittest
try:
    from ..binary_tree import in_order, Node
except:
    from ds.tree.binary_tree import in_order, Node


class TestBinaryTree(unittest.TestCase):
    # override in-built method setUp, to setup resources
    def setUp(self):
        self.root = Node(1)
        self.root.left = Node(2)
        self.root.right = Node(3)
        self.root.left.left = Node(4)
        self.root.left.right = Node(5)
        self.result = []
        self.func_in_order = in_order(self.root, self.result)

    def test_in_order_positive(self):
        self.assertEqual(self.func_in_order, [4, 2, 5, 1, 3], msg=None)

    def test_in_order_negative(self):
        self.assertNotEqual(self.func_in_order, [], msg=None)


if __name__ == '__main__':
    root = None
    unittest.main()
