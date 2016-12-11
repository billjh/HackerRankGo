"""
 Reverse a linked list
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the head of the linked list in the below method.
"""

def Reverse(head):
    # Return if the list has 0 or 1 Node
    if head == None or head.next == None:
        return head
    prev = None
    cur = head
    while cur.next != None:
        tmp = cur.next
        # Redirect cur.next to prev Node
        cur.next = prev
        # Move forward prev, cur Nodes
        prev = cur
        cur = tmp
    cur.next = prev
    return cur
