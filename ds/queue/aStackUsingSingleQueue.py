import queuelib
from queuelib.queue import FifoDiskQueue
class StackFromQueue:
    
    q = queuelib.queue.FifoDiskQueue("D:",100)    
    def __init__(self):
        #self.q.pop()
        pass
    def enqueue(self,val):
        self.q.push(val)
    def dequeue(self):
        first=self.q.pop()
        return first
    
q1=StackFromQueue()
q1.enqueue()
print(q1.dequeue())