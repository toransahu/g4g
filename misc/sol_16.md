## 16. Design an efficient data structure that can do following operations:

- findMin() returns minimum item, frequency of operation: mostly
- findMax() return maximum item, frequency : mostly
- deleteMin(): delete Minimum item, frequency: moderate
- deleteMax: delete maximum item, frequency: moderate,
- insert: insert an item, frequency: least
- delete: delete an item, frequency: least

Hmmm.. find/delete Min/Max resembles with min/max heap. If we maintain 1 min heap and 1 max heap and insert every item to each heap, we're close.

The only thing is to synchronize both the heaps on insert, delete, delete min/max.

### Approach 1

- 1 min heap (array): nodes stores the data
- 1 max heap (array): nodes stores the data
- 1 map: key=data, value=index of the data in min heap array
- 1 map: key=data, value=index of the data in max heap array

On perculation of node in heap, update the indices of the nodes in the maps.

Space complexity: O(c \* n) == O(n)

Time complexity:

- findMin: O(1)
- findMax: O(1)
- deleteMin: O(log n)
- deleteMax: O(log n)
- insert: O(log n), worst O(n)
- delete: O(log n)

Note: Not a stable algorithm, as the position/order of the items could not be preserved.

### Approach 2

- 1 min heap (array): nodes stores the address of a binary node
- 1 max heap (array): nodes stores the address of a binary node
- n binary nodes: node_data=item, left=index of the item in min heap array, right=index of the item in max heap array

Space complexity: O(c * n) == O(n)

Time complexity:

- findMin: O(1)
- findMax: O(1)
- deleteMin: O(log n)
- deleteMax: O(log n)
- insert: O(log n)
- delete: O(log n)
