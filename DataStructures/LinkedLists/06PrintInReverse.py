"""
 Print elements of a linked list in reverse order as standard output
 head could be None as well for empty list
 Node is defined as

 class Node(object):

   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node


"""

def ReversePrint(head):
    if head == None:
        return
    l = [head.data]
    cur = head
    while cur.next != None:
        # Push Node data into the start of a list
        cur = cur.next
        l = [cur.data] + l
    for i in l:
        # Print data by order
        print(i)
