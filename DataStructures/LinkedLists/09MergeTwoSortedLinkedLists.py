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
    # Return list B if A is empty
    if headA == None:
        return headB
    # Return List A if B is empty
    if headB == None:
        return headA
    # Mark the start of merged list
    if headA.data < headB.data:
        start = headA
        headA = headA.next
    else:
        start = headB
        headB = headB.next
    cur = start
    while headA != None and headB != None:
        # Keep merging when both remaining lists are not empty
        if headA.data < headB.data:
            cur.next = headA
            headA = headA.next
        else:
            cur.next = headB
            headB = headB.next
        cur = cur.next
    # Attach remaining list B if remaining list A is empty
    if headA == None:
        cur.next = headB
    # Attach remaining list A if remaining list B is empty
    if headB == None:
        cur.next = headA
    # Return the start of merged list
    return start
