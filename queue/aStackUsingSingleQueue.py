
from collections import deque

class StackFromQueue:
    q=deque()
    def __init__(self):
        self.q.pop()
    def enqueue(self,val):
        self.q.append(val)
        return True
    def deque(self):
        top=self.q.popfromleft()
        return top()
sfq=StackFromQueue()
sfq.enqueue('a')
print(sfq)
 
