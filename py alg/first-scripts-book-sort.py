def bin_search(arr,item): 
    low = 0 
    max = len(arr) - 1 

    while low <= max:
        mid = (low + max) // 2 
        guess = arr[mid]
        if guess == item:
            return mid 
        elif guess > mid:
            max = mid -1 
        else:
            low = mid + 1 
    return None
my_list = [1,2,3,4,5,6]
print(bin_search(my_list,2))