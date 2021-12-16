from queue import Queue
import math

class TreeNode:
    def __init__(self,val=None):
        self.val = val
        self.left = None
        self.right = None
        self.parent = None

class BinaryTreeNode:
    def __init__(self,val_list=[]):
        self.root = None
        for n in val_list:
            self.insert(n)
    
    def insert(self,data):
        if self.root is None:
            self.root =TreeNode(data)
        else:
            n=self.root
            while n:
                p=n
                if data < n.val:
                    n=n.left
                else:
                    n=n.right
            new_node=TreeNode(data)
            new_node.parent=p
            if data<p.val:
                p.left=new_node
            else:
                p.right=new_node
        return True

    def search(self,data):
        ret=[]
        n=self.root
        while n:
            if data<n.val:
                n=n.left
            else:
                if data==n.val:
                    ret.append(n)
                n=n.right
        return ret

    def delete(self,data):
        del_list=self.search(data)
        for n in del_list:
            if n.parent is None and n!=self.root:
                continue
            else:
                self.deleteNode(n)

    def deleteNode(self,node):
        if node.left is None and node.right is None:
            if node==self.root:
                self.root=None
            else:
                if node.val<node.parent.val:
                    node.parent.left=None
                else:
                    node.parent.right=None
                node.parent=None
        elif node.left is None and node.right is not None:
            if node==self.root:
                self.root=node.right
                self.root.parent=None
                node.right=None
            else:
                if node.val < node.parent.val:
                    node.parent.left= node.right
                else:
                    node.parent.right=node.right
                node.p