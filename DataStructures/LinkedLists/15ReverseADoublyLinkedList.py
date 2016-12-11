"""
 Reverse a doubly linked list
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None, prev_node = None):
       self.data = data
       self.next = next_node
       self.prev = prev_node

 return the head node of the updated list
"""
def Reverse(head):
    cur = head
    while cur != None:
        # Swap Node.next and Node.prev pointers
        cur.next, cur.prev = cur.prev, cur.next
        if cur.prev == None:
            # Until reach the end of the list
            return cur
        cur = cur.prev
