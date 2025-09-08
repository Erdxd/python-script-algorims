def bin_search(arr,item): 
    low = 0 
    maxi = len(arr) - 1 

    while low <= maxi:
        mid1: int = (low + maxi) // 2 
        guess = arr[mid1]
        if guess == item:
            return mid1 
        elif guess > item:
            maxi = mid1 -1 
        elif guess < item:
            low = mid1 + 1 
    return None 
my_list = [1,2,3,100,6,7,8]
print(my_list)
my_list.sort()
print(my_list)
print(bin_search(my_list,6))