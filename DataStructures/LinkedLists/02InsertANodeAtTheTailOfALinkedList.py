"""
 Insert Node at the end of a linked list
 head pointer input could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the head of the linked list in the below method
"""

def Insert(head, data):
    if head == None:
        return Node(data, None)
    cur = head
    while cur.next != None:
        cur = cur.next
    cur.next = Node(data, None)
    return head
