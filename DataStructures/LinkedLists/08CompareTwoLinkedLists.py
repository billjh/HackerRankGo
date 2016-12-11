"""
 Compare two linked list
 the lists are equal only if they have the same number of nodes
 and corresponding nodes contain the same data
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return 1 if the lists are equal; otherwise, return 0
"""

def CompareLists(headA, headB):
    # Traverse both lists Node by Node
    while headA != None and headB != None and headA.data == headB.data:
        # Don't stop only when both have equal data
        headA, headB = headA.next, headB.next
    if headA == None and headB == None:
        # Reach the end the two lists: equal
        return 1
    # Otherwise: not equal
    return 0
