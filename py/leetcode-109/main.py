from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def sortedListToBST(self, head: Optional[ListNode]) -> Optional[TreeNode]:
        # divide and conquer
        # use middle node as root
        # then convert left half of list to BST and right half to BST

        if head is None:
            return None

        n = 0
        curr = head
        while curr is not None:
            n += 1
            curr = curr.next

        return self.rec(head, n)

    def rec(self, head: Optional[ListNode], n: int) -> Optional[TreeNode]:
        if n == 1:
            return TreeNode(val=head.val)
        if n == 2:
            return TreeNode(val=head.val, right=TreeNode(val=head.next.val))

        mid = n//2
        curr = head
        i = 0
        while curr is not None:
            if i == mid:
                break

            i += 1
            curr = curr.next

        left = head
        root = curr
        right = curr.next

        # e.g. n = 5, mid=2, left are indices 0,1  right 3,4
        return TreeNode(val=root.val, left=self.rec(left, mid), right=self.rec(right, n-mid-1))
