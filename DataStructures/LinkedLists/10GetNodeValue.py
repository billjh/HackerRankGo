"""
 Get Node data of the Nth Node from the end.
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the node data of the linked list in the below method.
"""

def GetNode(head, position):
    if head == None:
        return None
    # Actually it's (postion + 1)-th Node from the end
    pos, count = position + 1, 0
    # Cache data of last n Nodes
    cache = [None] * pos
    while head != None:
        # Move towards the end of list and refresh cache
        cache[count%pos] = head.data
        count += 1
        head = head.next
    # Return data from cache
    return cache[count%pos]
