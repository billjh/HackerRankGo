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
    pos = position + 1
    count = 0
    cache = [None] * pos # an array to cache position + 1 node data
    while head.next != None:
        cache[count%pos] = head.data
        count += 1
        head = head.next
    cache[count%pos] = head.data
    count += 1
    return cache[count%pos]
