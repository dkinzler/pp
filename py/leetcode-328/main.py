from typing import Optional

#Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return head

        odd_head, even_head = None, None
        odd, even = None, None
        i = 1
        curr = head
        while curr is not None:
            next = curr.next
            curr.next = None
            if i % 2 == 1:
                if odd is None:
                    odd = curr
                    odd_head = curr
                else:
                    odd.next = curr
                    odd = curr
            else: 
                if even is None:
                    even = curr
                    even_head = curr
                else:
                    even.next = curr
                    even = curr

            i += 1
            curr = next

        odd.next = even_head
        return odd_head