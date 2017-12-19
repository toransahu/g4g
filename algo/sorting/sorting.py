#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Mon Dec 11 09:03:20 2017

@author: toran
"""


def heapify_subtree(arr, size, node):
    """
    Heapify a node of a Sub Tree of a given size.
    """
    n = size
    # assign i to largest
    largest = node
    # calculate index of left and right_child child
    left_child = 2 * node + 1
    right_child = 2 * node + 2
    
    # check if left child exists and
    # left child is larger than its parent
    if left_child < n and arr[largest] < arr[left_child]:
        largest = left_child
    
    # check if right child exists and
    # right child is larger than its parent
    if right_child < n and arr[largest] < arr[right_child]:
        largest = right_child
    
    # if the root/parent has been changed (from its orignal/initial value),
    # i.e. the index of largest and node is not same as we had assigned earlier,then
    # swap the value of larger child with parent
    if largest != node:
        arr[largest], arr[node] = arr[node], arr[largest]
        # and heapify that node      i.e. the index of child which was larger
        heapify_subtree(arr, size, largest)            
    return arr

def heap_sort(arr, left, right):
    """
    Heap Sort.
    
    Desc:
    """
    pass
    

def partition(arr, left, right):
    """
    Partition.
    
    Picks pivot is any preferable fashion:
        1. First element
        2. Last element
        3. Mean
        4. Random
    Finds index of pivot and re-arranges elements smaller than pivot to left
    and greater to right.
    Returns: Partitioning Index (Re-calculated index of pivot)
    Time Complexity: O(n)
    Auxilary Space: O(1)
    """
    pivot_idx = right
    pivot_val = arr[pivot_idx]
    partition_idx = left
    
    # iterate from left to one step less of right
    # check number of elements lower than pivot_val, it will give partition_idx
    
    for i in range(left,right):
        if arr[i]<=pivot_val:
            arr[i], arr[partition_idx] = arr[partition_idx], arr[i]
            partition_idx+=1
    
    # put pivot in its partitioning index
    arr[right], arr[partition_idx] = arr[partition_idx], arr[right]
    return partition_idx

def quick_sort_iterative(arr,left,right):
    """
    Desc:
        In this version of quicksort we're trying to mimic "how compiler executes 
        recursive functions using recursion stacks".
        The performance of this version may be worst than recursive version.
        
    Auxilary Space: O(n)
    """
    
    # create an auxilary stack of any size, does not matter
    # we will use top to navigate
    # DO NOT USE inbuit Methods, may increase the time complexity
    top = -1
    stack = [0]*(right - left + 1)
    top+=1
    stack[top] = left
    top+=1
    stack[top] = right
    
    # keep popping from stack while its not empty
    # do partition
    # re-calculate left & right in binary fashion
    while(top>=0):
        # pop left & right
        right = stack[top]
        top-=1
        left = stack[top]
        top-=1
        
        # partition
        partition_idx = partition(arr, left, right)
        
        # check if there is any sub-array in LHS of partition_idx
        if left < partition_idx-1:
            # fill the stack with left & right of LHS sub-array
            top+=1
            stack[top] = left
            top+=1
            stack[top] = partition_idx - 1
            
        # check if there is any sub-array in RHS of partition_idx
        if partition_idx + 1 < right:
            # fill the stack with left & right of RHS sub-array
            top+=1
            stack[top] = partition_idx + 1
            top+=1
            stack[top] = right
    return arr
            

def quick_sort_recursive(arr,left,right):
    """
    Quick Sort.
    
    * Description:
            1. Partition the array about a pivot: re-arrange smaller elements in LHS & greater elements in RHS of pivot
            2. Return partitioned index (i.e. index of pivot after re-arrangement)
            3. Re-call quicksort for sub-array smaller than pivot
            4. Re-call quicksort for sub-array greater than pivot
        * Time Complexity:
            1. Best Case: O(nlogn) (Occurs when pivot element is middle element value-wise)
            2. Avg Case: O(nlogn) 
            3. Worst Case: O(n2) (Occurs when pivot element is either smaller or larger element)
        * Auxilary Space: O(1)
        * Algorithm Paradigm: Divide and Conquer
        * Implementation: Recursive (Generally) and Iterative
        * In-Place: Yes (because auxilary space O(n)) 
    """
    if left<right:
        partition_idx = partition(arr,left,right)
        quick_sort_recursive(arr,left,partition_idx-1)
        quick_sort_recursive(arr,partition_idx+1,right)
    return arr


def merge(arr,left, pivot, right):
    """
    Merge.
    
    Desc:
        1. Merges two sorted sub-arrays by using an extra space of O(n).
        2. begin from 0th index of both sub-array, do comparision
        3. Make changes in original array
    Time Complexity: O(n)
    Auxilary Space: O(n)
    """
    import copy
    i = left
    j = pivot+1
    k = left
    temp_arr = copy.copy(arr)
    while(i<=pivot and j<=right):
        if temp_arr[i] <= temp_arr[j]:
            arr[k] = temp_arr[i]
            i+=1
        else:
            arr[k] = temp_arr[j]
            j+=1        
        k+=1
    while(i<=pivot):
        arr[k] = temp_arr[i]
        i+=1
        k+=1
    while(j<=right):
        arr[k] = temp_arr[j]
        j+=1
        k+=1
    return arr
        

def merge_sort(arr, left, right):
    """
    Merge Sort.
    
    Description: 
        1. Binary split of array till minimum/single entity
        2. Start merging single entity --> 2 sorted sub-array --> merge both
        3. Continue merging till end
    Paradigm: Divide  & Conquer
    Implementation: Recursive only
    In-place: Yes
    Time Complexity: 
        Best Case: O(nlogn)
        Avg Case: O(nlogn)
        Worst Case: O(nlogn)
    Auxilary Space: O(n)
    """
    if(left<right):
        pivot = (left+right)//2
        merge_sort(arr,0,pivot)
        merge_sort(arr,pivot+1,right)
        merge(arr,left,pivot,right)
    return arr
    
       
        
def insertion_sort(arr):
    """
    Select second & onwards elements one by one. ie. i
    Compare rest elements arr[j] i.e. 0...i-1 with above elements & if arr[j]
    is greater than arr[i] then shift jth elements right by one index.
    Repeat this for all j from i-1 to 0.
    
    Time Complexity: O(n2)
    Auxilary Space: 
    """
  
    n = len(arr)
    for i in range(1,n):
        k = i-1
        v = arr[i]
        for j in range(i-1,-1,-1):            
            if arr[j]>v:
                arr[j+1] = arr[j]
                k = j
        arr[k] = v        
    return arr
    

def selection_sort(arr):
    """
    Pulls a minimum element from sub-array and appends infront of that array.
    Time Complexity: 
        all case: O(n2)
    Auxilary Space: O(1)
    """
    n = len(arr)
    for i in range(0,n-1):
        min_idx = i
        for j in range(i,n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        arr[i], arr[min_idx] = arr[min_idx], arr[i]
    return arr

def bubble_sort_optimized(arr):
    """
    Optimized to detect an already sorted array.
    Time Complexity:
        Worst/Avg Case: O(n2)
        Best Case: O(n)
    """
    n = len(arr)
    swapped = False
    # travers the array n-1 times
    for i in range(0,n-1):
        # Last i elements are already in place
        for j in range(0,n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = True
        if not swapped:
            return arr
    return arr


def bubble_sort(arr):
    """
    Time Complexity:
        Worst/Avg Case: O(n2)
        Best Case: O(n2)
    """
    n = len(arr)
    # travers the array n-1 times
    for i in range(0,n-1):
        # Last i elements are already in place
        for j in range(0,n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
    return arr


arr = [3,6,9,1,4,9,0,3,5,2]
arr = [9,8,7,6,5,4,3,2,1]
arr_s = [1,2,3,4,5,5,6,7]
# print(bubble_sort(arr))
# print(bubble_sort_optimized(arr_s))
# print(selection_sort(arr))
# print(insertion_sort(arr))
# print(insertionSort(arr))
# print(merge_sort(arr,0,8))
# print(quick_sort_recursive([2,3,9,4,5,10],0,5))
# print(quick_sort_iterative([2,3,9,4,5,10],0,5))
# print(quick_sort_recursive([1,1,1,1,1,1],0,5))
# print(max_heap([1,3,2,5,4,0],0,5))
print(max_heap([1,2,3,4,5,6],0))