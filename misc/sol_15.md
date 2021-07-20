## 15. design a data structure that can do following operations in O(1) time

Insert(X), delete(X), search (X) and getRandom()

This mostly resembles with a Map/Dict data structure where

- delete(num) is O(1)
- search(num) is O(1)
- insert(num) is O(1) and O(n) in worst case

The only thing left is getRandom(). In Map/Dict we can't access first/last or any key-val by some fixed+known address/index. If this would have known, then we could have pick some (or same) address/index and respond with corresponding value. But, unfortunately, we cannot go this way.

So, lets maintain a List/Dynamic array along with a Map/dict. The list will store the keys of the Map. We need to make sure we omit out the deleted (flagged deleted/false/null) items.
