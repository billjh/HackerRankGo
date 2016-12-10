"""
 Merge two linked lists
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the head of the linked list in the below method.
"""

def MergeLists(headA, headB):
    if headA == None:
        return headB
    if headB == None:
        return headA
    if headA.data < headB.data:
        start = headA
        headA = headA.next
    else:
        start = headB
        headB = headB.next
    cur = start
    while headA != None and headB != None:
        if headA.data < headB.data:
            cur.next = headA
            headA = headA.next
        else:
            cur.next = headB
            headB = headB.next
        cur = cur.next
    if headA == None:
        cur.next = headB
    if headB == None:
        cur.next = headA
    return start
