"""
 Delete duplicate nodes
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the head of the linked list in the below method.
"""

def RemoveDuplicates(head):
    if head == None:
        return head
    # Cache the current Node and data
    cur, value = head, head.data
    while cur.next != None:
        if value != cur.next.data:
            # Refresh cache
            cur, value = cur.next, cur.next.data
        else:
            # Skip the Node by redirecting next pointer
            cur.next = cur.next.next
    return head
