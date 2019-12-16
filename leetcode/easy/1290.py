# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getDecimalValue(self, head: ListNode) -> int:
        arr = [];
        while (head.next != None):
            arr.append(head.val);
            head = head.next;
         
        arr.append(head.val);
        arr = arr[::-1];
        res = 0;
        
        for i in range(len(arr)):
            res += arr[i] * pow(2, i);
        
        return res;

## ans = 0
##   while head:
##     ans = (ans << 1) | head.val
##     head = head.next
##     return ans
