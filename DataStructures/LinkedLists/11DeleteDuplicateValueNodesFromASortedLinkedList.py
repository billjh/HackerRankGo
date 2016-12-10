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
    cur, value = head, head.data
    while cur.next != None:
        if value != cur.next.data:
            cur, value = cur.next, cur.next.data
        else:
            cur.next = cur.next.next
    return head
