<<<<<<< HEAD
'''
Created on 10-Sep-2016

@author: toran
'''
import sys
import queue
import queue.Queue

class A:Queue
    pass


        
=======
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
 
        

>>>>>>> 4e4ce432534510ab0c8650b0c95efdd83b610319
