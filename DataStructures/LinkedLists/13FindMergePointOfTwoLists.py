"""
 Find the node at which both lists merge and return the data of that node.
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node
"""

def FindMergeNode(headA, headB):
    a = set()
    while headA != None:
        # Put all Nodes (object references) from list A into a set
        a.add(headA)
        headA = headA.next
    while headB != None:
        # Search for the merge Node from list B
        if headB in a:
            return headB.data
        headB = headB.next
    return None
