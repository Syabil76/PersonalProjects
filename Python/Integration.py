#to be paired with my functions code

import math
import string

f_array = []
funny = []
degree = int(input("input degree of polynomial"))

def f():
    if degree > 4:
        print("this is above my pay grade")
    print("input function")
    for i in range(degree,-1,-1):
        if i == 0:
            constant = int(input(f"input '{chr(97+degree)}' value"))
            f_array.append(constant)
            funny.append(constant)
        elif i == 1:
            linear = int(input(f"input '{chr(97+degree-1)}x' value"))
            f_array.append(str(linear)+"x+")
            funny.append(linear)
        else:
            polynomial = int(input(f"input '{chr(97+degree-i)}x^{i}' value"))
            f_array.append(str(polynomial)+f"x^{i}+")
            funny.append(polynomial)
        print(f_array)
    function = ''.join(map(str, f_array))
    print(function)
    print(funny)

def integration():
    integrated = []
    newpoly = []
    for i in range(degree+1, 0, -1):
        print(f"{funny[degree+1-i]} to the power of {i} divided by {i}")
        new = round(
            ((funny[degree+1-i])/i), 5
        )
        integrated.append(new)
        newpoly.append(f"{str(new)}x^{i}+")
    print(f"indefinite integral = {''.join(map(str,newpoly))} C")
    q = input("do you want to find area? y/n")
    if q == 'y':
        lower = int(input("what is your lower bound?"))
        upper = int(input("what is your upper bound?"))
        uppertotal = 0
        lowertotal = 0
        total = 0
        for j in range(len(integrated)):
            print(integrated[j])
            uppertotal = uppertotal+(upper**(degree+1-j)*integrated[j])
            lowertotal = lowertotal+(lower**(degree+1-j)*integrated[j])
        total = uppertotal-lowertotal
    else:
        return
    print(f"total is {total}")
f()
integration()

