# 2.Add Two Numbers
# You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit.
# Add the two numbers and return it as a linked list.
# 
# You may assume the two numbers do not contain any leading zero, except the number 0 itself.
#
## Example:
#
# Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
# Output: 7 -> 0 -> 8
# Explanation: 342 + 465 = 807.
####

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        guard = ListNode(None)
        node_a, node_b, current, rest = l1, l2, guard, 0
        
        while node_a or node_b or rest:
            total = (node_a.val if node_a else 0) + (node_b.val if node_b else 0) + rest
            
            current.next = ListNode(total % 10)
            current = current.next
            rest = total // 10
            
            node_a = node_a.next if node_a else None
            node_b = node_b.next if node_b else None

        return guard.next
