def fact(a):
    z = 0 
    p = 1
    while z < a:
        z+= 1
        p *= z
    return p 
print(fact(11))
print(fact(2))
        
