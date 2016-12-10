"""
 Delete Node at a given position in a linked list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the head of the linked list in the below method.
"""

def Delete(head, position):
    if head == None or head.next == None:
        return None
    if position == 0:
        return head.next
    cur = head
    while cur.next != None and position > 0:
        prev = cur
        cur = cur.next
        position -= 1
    prev.next = cur.next
    return head
