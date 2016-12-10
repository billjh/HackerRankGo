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
    if head == None or head.next == None:
        return head
    prev = None
    cur = head
    while cur.next != None:
        tmp = cur.next
        cur.next = prev
        prev = cur
        cur = tmp
    cur.next = prev
    return cur
