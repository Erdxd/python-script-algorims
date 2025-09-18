def sumwithrecursia(arr):
    if arr == []:
        return 0 
    return arr[0] + sumwithrecursia(arr[1:])
        
d = [1,2,3]  
print(sumwithrecursia(d))