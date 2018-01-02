#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Mon Dec 11 09:03:20 2017

@author: toran
"""


def shell_sort(arr):
    """
    Shell Sort.

    * **Pre-requisites:**
    * **Desc:**
    * **Useful:**
    * **Advantage:**
    * **Applications:**
    * **Recurrence Equation:**
    * **Time Complexity:**
    * **Auxilary Space:**
    * **In-Place:**
    * **Implementation:**
    * **Algorithm Paradigm:**
    * **Data Structure:**
    * **Stable:**
    * **Note:**
    """
    pass


def counting_sort_for_radix(arr, decimal_place):
    """
    Modified Counting Sort as Sub-Routine in Radix Sort.

    * **Pre-requisites:**
    * **Desc:**
    """
    size = len(arr)
    # create counting array of length [0 to 9]
    # such that we can store counts of digit at current decimal place
    # auxilary space:
    counts = [0] * 10

    # create an extra output array of size same as input array
    # auxilary space: O(n)
    output = [0] * size

    # start storing counts of each digit in count array
    # time complexity: O(n)
    for i in arr:
        current_index = i // decimal_place
        current_digit = current_index % 10
        counts[current_digit] += 1

    # add previous count to next count
    # so that each count can represent the last occurance index of each digit
    # time complexity: O(n)
    for i in range(1, 10):  # 10 == len(counts)
        counts[i] += counts[i - 1]

    # REVERSE iterate over input array (TO MAKE IT STABLE SORT) and find position of current digit in count array
    # store the original element in output array and decrease the count by 1
    # time complexity: O(n)
    for i in arr[::-1]:
        current_index = i // decimal_place
        current_digit = current_index % 10
        output[counts[current_digit] - 1] = i
        counts[current_digit] -= 1

    return output


def radix_sort(arr):
    """
    Radix Sort.

    * **Pre-requisites (Standard):**
        * input array have non negative integers
        * range should be 0 to n**c where c is some constant & numbers are represented in base n
        * or each number takes only log2(n) bits
    * **Desc:**
        * for 1 to d: where d is most significant digit position of MAX element in inp array
        * do counting sort on array
    * **Useful:** same as prerequisites
    * **Advantage:** Over bucket sort Worst case
    * **Applications:** Card sorting machine
    * **Recurrence Equation:** n*O(n + k) == O(n + k)
    * **Time Complexity:**
        * Best Case: Omega(n + k)
        * Avg Case: theta(n + k)
        * Worst Case: O(n + k)
    * **Auxilary Space:** d*O(counting array + output array) = d*O(n + k)
    * **In-Place:** No
    * **Implementation:** Iterative
    * **Algorithm Paradigm:** Partial Hashing
    * **Data Structure:** Hashtable, array
    * **Stable:** Yes
    * **Comparion Sort:** No
    * **Note:**
    """

    # Find MAX element in the input array
    MAX = arr[0]
    for element in arr:
        if MAX < element:
            MAX = element

    # find d=number of digits of MAX element
    d = 1
    while MAX > 0:
        MAX = MAX // 10
        d += 1

    # store decimal places possible in input array to decimal_places array
    decimal_places = [10**i for i in range(0, d - 1)]
    # do counting sort d times == len(decimal_places)
    for decimal_place in decimal_places:
        arr = counting_sort_for_radix(arr, decimal_place)
        print(arr)
    return arr


def counting_sort(arr):
    """
    Counting Sort

    Pre-requisites (Standard):
        1. Elements should be Non Negative Integers
        2. Over a range of 0 to k where k < size of array to maintain O(n)
    Desc:
        For each element X in the input array find the number of elements smaller than X.
    Steps:
        1. Store counts of each element in a counting array
        2. Add previous count to current count, to find index of last occurence of that element
        3. REVERSE Iterate over input array & pick index of the element from counting array
        4. Put the element in output array and decrement the count by 1
    Useful: same as pre-requisites
    Advantage:
        1. Sorting in O(n + k)
    Applications:
        1. Same as pre-requisites
        2. As a subroutine in Radix Sort
    Time Complexity:
        Best: Omega(n + k)
        Avg: Theta(n + k)
        Worst: O(n + k)
    Auxilary Space: O(counting + output array) == O(n + k)
    In-Place: No
    Implementation: Iterative
    Algorithm Paradigm: Partial Hashing
    Data Structure: Hashtable, Array
    Stable: Yes (order of elements with same value in input array maintains same order in output)
    Comparion Sort: No
    Note: Can be extended to sort negative integers also
    """
    size = len(arr)
    # find max element (i.e. MAX will be k )
    MAX = arr[0]
    for i in arr:
        if i > MAX:
            MAX = i

    # create counting array of length = value of MAX element + 1
    # such that we can store counts of elements in a range (0 to MAX)
    # auxilary space: O(k)
    counts = [0] * (MAX + 1)

    # create an extra output array of size same as input array
    # auxilary space: O(n)
    output = [0] * size

    # start storing counts of each elements in count array
    # time complexity: O(n)
    for i in arr:
        counts[i] += 1

    # add previous count to next count
    # so that each count can represent the last occurance index of each element
    # time complexity: O(n)
    for i in range(1, MAX + 1):  # MAX + 1 == len(counts)
        counts[i] += counts[i - 1]

    # REVERSE iterate over input array (TO MAKE IT STABLE SORT) and find their position in count array
    # store the element in output array and decrease the count by 1
    # time complexity: O(n)
    for i in arr[::-1]:
        output[counts[i] - 1] = i
        counts[i] -= 1

    return output


def bucket_sort(arr):
    """
    Bucket Sort (Generalized).

    Pre-requisites:
    * In Comman: A uniform distributed input array in a range of [0,1).
    * Generalized: A uniform distributed input array in a range of non negative integers + floats.
    * Efficient Hash Function (specially in case of "Generalized" implementation.
    Desc:
        1. Hashing:
            hash_table_size or number of buckets:
                * = size of input array; Standard
                * OR = sqrt(size); Generalized
            hash_func() = (element/MAX)* (hash_table_size)
            Condition: if i < k then hash(i) < hash(k)
        2. Partion inp array on the basis of hash function, store then in right bucket/array.
        3. Sort each array  using Insertionsort
        4. Merge all sorted arrays into one.
    Useful: When input is uniformly distributed over a positive range
    Advantage:
        1. Sorting in O(n)
    Applications:
        1. When input is uniformly distributed over a positive range
    Recurrence Equation: Theta(n) + n.O(2 - 1/n)
    Time Complexity:
        Best: Omega(n); If each bucket-array have 1 element
        Avg: Theta(n)
        Worst: Theta(n2); If all elements falls under single bucket
    Auxilary Space: O(bucket size) == O(n)
    In-Place: No
    Implementation: Iterative
    Algorithm Paradigm: Hashing, Partion
    Data Structure: Hashtable, Array
    Stable: Yes
    Note: If input is not uniformally distributed, but also bucketsort may still run in linear time
    """
    size = len(arr)

    # find max element in array
    # Time Complexity: O(n)
    MAX = arr[0]
    for i in arr:
        if MAX < i:
            MAX = i

    # Choose size of hashtable (or number of buckets) & hash function
    import math

    hash_table_size = int(math.sqrt(size))
    # hash_table_size = size
    hash_func = lambda element: int((element / MAX) * (hash_table_size - 1))

    # create "hash_table_size" number of arrays/buckets
    # Auxilary Space: O(n) in worst case
    buckets = [list() for _ in range(hash_table_size)]

    # iterate through all the elements and
    # find hash value for that element and
    # store that element in right bucket, according to hash value
    # Time Complexity: O(n)
    for i in arr:
        hash_value = hash_func(i)
        buckets[hash_value].append(i)

    # sort all the arrays/buckets in the hashtable
    # Time Complexity: n*O(2 - 1/n) #HOW???? #READ NOTES or CLRS Book
    for array in buckets:
        insertion_sort(array)

    # merge all the sorted arrays
    # time complexity: O(n)
    idx = 0
    for array in buckets:
        for element in array:
            arr[idx] = element
            idx += 1

    return arr


def heapify_subtree(arr, size, node):
    """
    Heapify a node of a Sub Tree of a given size.

    Time Complexity: O(logn)
    Auxilary Space: O(1)
    Implementation: Recursive
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


def heap_sort(arr):
    """
    Heap Sort.

    Desc:
        1. Make max heap using heapify(). To do so, start from last leaf and make max heap till root node.
        2. Loop from last leaf to second node.
            1. Swap last leaf with root node in max heap
            2. Heapify the sub-tree (by ignoring last leaf) at root node
               till the loop covers all the nodes except the root node
    Useful:
        1. When there is time (Quick's problem) and space (Merge's problem) bound
    Advantage:
        1. worst case upper bound is O(nlogn) with only O(1) auxilary space
    Applications:
        1. Sort a nearly sorted (or K sorted) array
        2. k largest(or smallest) elements in an array
    Time Complexity: O(heapify) * O(n) = O(nlogn)
    Auxilary Space: O(1)
    In-Place: Yes
    Implementation: Recursive (Heapify)
    Algorithm Paradigm: Comparision Based
    Data Structure: Array + Complete Binary Tree
    Stable: Not in general
    Note: Quick & Merge are better in practice.
    """
    size = len(arr)
    # make max heap
    # to do so, start from last leaf and make max heap till root node
    for node in range(size, -1, -1):
        heapify_subtree(arr, size, node)

    # swap last leaf with root node in max heap
    # and heapify the sub-tree (by ignoring last leaf) at root node
    # till the loop covers all the nodes except the root node
    for i in range(len(arr) - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapify_subtree(arr, i, 0)
    return arr


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

    for i in range(left, right):
        if arr[i] <= pivot_val:
            arr[i], arr[partition_idx] = arr[partition_idx], arr[i]
            partition_idx += 1

    # put pivot in its partitioning index
    arr[right], arr[partition_idx] = arr[partition_idx], arr[right]
    return partition_idx


def quick_sort_iterative(arr, left, right):
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
    stack = [0] * (right - left + 1)
    top += 1
    stack[top] = left
    top += 1
    stack[top] = right

    # keep popping from stack while its not empty
    # do partition
    # re-calculate left & right in binary fashion
    while (top >= 0):
        # pop left & right
        right = stack[top]
        top -= 1
        left = stack[top]
        top -= 1

        # partition
        partition_idx = partition(arr, left, right)

        # check if there is any sub-array in LHS of partition_idx
        if left < partition_idx - 1:
            # fill the stack with left & right of LHS sub-array
            top += 1
            stack[top] = left
            top += 1
            stack[top] = partition_idx - 1

        # check if there is any sub-array in RHS of partition_idx
        if partition_idx + 1 < right:
            # fill the stack with left & right of RHS sub-array
            top += 1
            stack[top] = partition_idx + 1
            top += 1
            stack[top] = right
    return arr


def quick_sort_recursive(arr, left, right):
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
    if left < right:
        partition_idx = partition(arr, left, right)
        quick_sort_recursive(arr, left, partition_idx - 1)
        quick_sort_recursive(arr, partition_idx + 1, right)
    return arr


def merge(arr, left, pivot, right):
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
    j = pivot + 1
    k = left
    temp_arr = copy.copy(arr)
    while (i <= pivot and j <= right):
        if temp_arr[i] <= temp_arr[j]:
            arr[k] = temp_arr[i]
            i += 1
        else:
            arr[k] = temp_arr[j]
            j += 1
        k += 1
    while (i <= pivot):
        arr[k] = temp_arr[i]
        i += 1
        k += 1
    while (j <= right):
        arr[k] = temp_arr[j]
        j += 1
        k += 1
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
    if (left < right):
        pivot = (left + right) // 2
        merge_sort(arr, 0, pivot)
        merge_sort(arr, pivot + 1, right)
        merge(arr, left, pivot, right)
    return arr


def insertion_sort(arr):
    """
    Insertion Sort.

    Desc:
        Select second & onwards elements one by one. ie. target = arr[i]
        Compare rest elements arr[j] i.e. i-1...0 with above elements & if arr[j]
        is greater than arr[i] then shift jth elements right by one index.
        Repeat this for all j from i-1 to 0.
    Useful: For small or nearly sorted array.
    Time Complexity:
        Best: O(n); When array is already sorted
        Avg: O(n2);
        Worst: O(n2)
    Auxilary Space:
    """

    n = len(arr)
    # start looping from second element in the array
    for idx in range(1, n):
        # set current index value as target to put it at correct position
        target = arr[idx]
        # set pointer at index of the target element
        pointer = idx
        """
        #While way starts

        # loop if our pointer is not reached second element of the array and
        # our target is still smaller than its previous element
        while(pointer > 0 and target < arr[pointer - 1]):
            arr[pointer] = arr[pointer - 1]
            # shift pointer to left by one index
            pointer -= 1
        # put the target element at pointer position
        arr[pointer] = target

        #While way ends
        """

        for j in range(idx - 1, -1, -1):
            if target < arr[j]:
                arr[j + 1] = arr[j]
                pointer = j
            # This is necessary condition in for loop to achieve best case complexity
            else:
                break
        # put the target element at pointer position
        arr[pointer] = target
    return arr


def selection_sort(arr):
    """
    Pulls a minimum element from sub-array and appends infront of that array.
    Time Complexity:
        all case: O(n2)
    Auxilary Space: O(1)
    """
    n = len(arr)
    for i in range(0, n - 1):
        min_idx = i
        for j in range(i, n):
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
    for i in range(0, n - 1):
        # Last i elements are already in place
        for j in range(0, n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
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
    for i in range(0, n - 1):
        # Last i elements are already in place
        for j in range(0, n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
    return arr


arr = [3, 6, 9, 1, 4, 9, 0, 3, 5, 2, 20, 512, 312]
#arr = [9,8,7,6,5,4,3,2,1]
arr_s = [1, 2, 3, 4, 5, 5, 6, 7]
arr_float = [0.3, 0.1, 0.5, 0.2, 0.7, 0.9, 0.8]
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
# print(max_heap([1,2,3,4,5,6],0))
# print(heap_sort([1,2,3,4,5,6]))
# print(bucket_sort(arr_float + arr))
print(counting_sort(arr))
# print(radix_sort(arr))
#print(counting_sort_for_radix(arr, 10))
