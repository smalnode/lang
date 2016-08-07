class Node:
    def __init__(self, value):
        self._value = value
        self._left = None
        self._right = None

class BinaryTree:
    def __init__(self):
        self._root = None

    def insert(self, n):
        if not isinstance(n, Node):
            n = Node(n)
        if self._root is None:
            self._root = n
            return
        c = self._root
        p = None
        while c is not None:
            p = c
            if n._value < c._value:
                c = c._left
            else:
                c = c._right
        if n._value < p._value:
            p._left = n
        else:
            p._right = n

    def preorder(self):
        self._preorder(self._root)
        print()

    def _preorder(self, root):
        if root is None:
            return
        self._preorder(root._left)
        print(root._value, end=' ')
        self._preorder(root._right)
        
t = BinaryTree()
t.insert('K')
t.insert('A')
t.insert('C')
t.insert('E')
t.insert('R')
t.insert('M')
t.preorder()
