"""
Detect a cycle in a linked list. Note that the head pointer may be 'None' if the list is empty.

A Node is defined as:

    class Node(object):
        def __init__(self, data = None, next_node = None):
            self.data = data
            self.next = next_node
"""


def has_cycle(head):
    # Use Floyd's cycle-finding algorithm
    if head == None:
        return False
    # Traverse the list with two pointers of different speed: slow and fast
    slow = fast = head
    while fast != None and fast.next != None:
        slow = slow.next
        fast = fast.next.next
        # If there is a loop, fast will catch up slow
        if slow == fast:
            return True
    return False
