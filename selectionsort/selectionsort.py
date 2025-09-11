def findsmallest(arr):
    smallest = arr[0]
    smallest_index = 0 
    for i in range(1,len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i 
    return smallest_index
def selectionSort(arr):
    NewArr = []
    copiedArr = list(arr)
    for i in range(len(copiedArr)):
        smallest = findsmallest(copiedArr)
        NewArr.append(copiedArr.pop(smallest))
    return NewArr
arr = [1,2,100,9]
print(selectionSort(arr))
