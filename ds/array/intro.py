import array

# create
# arg1: typecode: 'i' for integer
# arg2: initializer: list of elements
arr = array.array("i", [1, 2, 3, 4])

# append
arr.append(5)

# extend: not available

# insert
# arg1: index
# arg2: element
arr.insert(2, 0)

# pop: removes element at index
# arg: index
# default: pops last element
# return: poped value
arr.pop()
arr.pop(2)


# remove: removes first occurance of the value
# arg: value
# return:
arr.remove(4)

# index: returns index of first occurance of the value
# arg: value
print(arr.index(3), "\n")

# reverse: permanently reverse the array
# return: None
print(arr.reverse(), "\n")

for i in range(0, len(arr)):
    print(arr[i])
