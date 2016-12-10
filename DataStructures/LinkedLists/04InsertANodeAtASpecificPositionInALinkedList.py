"""
 Insert Node at a specific position in a linked list
 head input could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the head of the linked list in the below method.
"""
#This is a "method-only" submission.
#You only need to complete this method.
def InsertNth(head, data, position):
    if head == None or position == 0:
        return Node(data, head)
    cur = head
    while cur.next != None and position > 0:
        prev = cur
        cur = cur.next
        position -= 1
    cur, cur.next = Node(data, cur), cur
    if head.next == None:
        head.next = cur
    else:
        prev.next = cur
    return head
