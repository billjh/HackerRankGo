"""
 Insert a node into a sorted doubly linked list
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None, prev_node = None):
       self.data = data
       self.next = next_node
       self.prev = prev_node

 return the head node of the updated list
"""
def SortedInsert(head, data):
    # Create the Node to be inserted
    newNode = Node(data, None, None)
    # Return the Node if list is empty
    if head == None:
        return newNode
    cur = head
    # Keep searching for the right position
    while cur.data <= data and cur.next != None:
        cur = cur.next
    if cur.data > data:
        # Case 1. Stop before a Node
        prev = cur.prev
        # Case 1a. Insert into the middle of list
        if prev != None:
            prev.next = newNode
            newNode.prev = prev
        # Case 1b. Insert into the start of list
        newNode.next = cur
        cur.prev = newNode
    else:
        # Case 2. Reach the end of the list
        cur.next = newNode
        newNode.prev = cur
    return head
