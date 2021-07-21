## 18. Suppose you are getting stream of words like army, Mary, tea, coffee, eat, kit, tik. Then you have to design a data structure to store these values such that if a query is made to search tik and it should also return it's anagram.

Write the function prototype as well.

## Solution

### Prototype

- myADT.Insert(givenString) -> boolean (true: if the givenString is inserted, false: if the given string already exits)
- myADT.Search(givenString) -> list of anagrams of givenString, including givenString (if any), else empty list


### Approaches

1. Try hash the input string such that all the anagrams also produces the same hash, then link all the given string with that hash
    - This could fail/perform bad if non-anagram string also produces the same hash
    - Time complexity
        - myADT.Insert(givenString): O(1 + m) --> O(n + m), where n is number of unique strings to store; m is number of anagrams of the givenString string
        - myADT.Search(givenString): O(1) --> O(n) where n is number of unique strings to store
    - Space complexity
1. Or, use the sorted string as a key/hash and link all the anagrams to that key
    - This could be implemented using:
        - A map, where key=sorted string, value=list of strings/anagrams
            - Time Complexity
                - myADT.Insert(givenString): O(n + m) where n is length of the givenString, and m is number of possible anagrams to be searched for duplicates
                - myADT.Search(givenString): O(n + m) where n is length of the givenString, and m is number of possible anagrams
                - where m=n! in worst case
                - So, time complexity could be O(n!) in worst case
                - We could improve upon this if we use Set or Map instead of an List/Array to store the anagrams
            - Space Complexity
                - O(n), where n is number of strings to store
        - Or, a n-ary Tree (say 26 children max), where given string is first sorted & then inserted to the tree. So, all the anagrams should lead to same leaf; where that leaf can store all those whole strings/anagrams
            - Time Complexity
                - myADT.Insert(givenString): O(n + m) where n is length of the givenString, and m is number of possible anagrams to be searched for duplicates
                - myADT.Search(givenString): O(n + m) where n is length of the givenString, and m is number of possible anagrams
            - Space Complexity
                - O(26^n * m), where n is max length of the input strings; m is number of strings to store
