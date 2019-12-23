# python Linked-List
# LC - 206
class ListNode:
    def __init__(self, x, next):
        self.val = x
        self.next = next

class Solution:
    # O(n)
    def fromArray(arr: list) -> ListNode:
        root = None
        for i in arr[::-1]:
            tmp = ListNode(arr[i], None)
            tmp.next = root
            root = tmp

        return root

    # O(n)
    def reverse(head: ListNode) -> ListNode:
        left = None
        while head:
            head.next, left, head = left, head, head.next

        head = left
        return left

    # O(n)
    def circle(head: ListNode) -> bool:
        while head.next.next:
            sPtr, fPtr = head.next, head.next.next
            head = head.next
            if sPtr.val == fPtr.val: return True
            

        return False
        

class Tests:
    def fa():
        l = [0, 1, 2, 3]
        ll = Solution.fromArray(l)
        ll2 = ListNode(0, ListNode(1, ListNode(2, ListNode(3, None))))
        while ll and ll2:
            assert ll.val is ll2.val
            ll, ll2 = ll.next, ll2.next

    def reverse():
        ll = ListNode(0, ListNode(1, ListNode(2, ListNode(3, None))))
        res = ListNode(3, ListNode(2, ListNode(1, ListNode(0, None))))
        ll2 = Solution.reverse(ll)

        while ll2.next and res.next:
            assert ll2.val is res.val
            ll2, res = ll2.next, res.next

    def circle():
        l = [0, 1, 2, 3]
        l2 = [0, 1, 2, 3, 3, 2]
        ll = Solution.fromArray(l)
        ll2 = Solution.fromArray(l2)

        assert Solution.circle(ll) is False
        assert Solution.circle(ll2) is True

    def run():
        Tests.fa()
        Tests.reverse()
        Tests.circle()

Tests.run()
        

