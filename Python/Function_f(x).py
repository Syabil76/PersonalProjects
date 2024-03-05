import math
import string


def f():
    f_array = []
    funny = []
    degree = int(input("input degree of polynomial"))
    if degree > 4:
        print("this is above my pay grade")
    print("input function")
    for i in range(degree,-1,-1):
        if i == 0:
            constant = int(input(f"input '{chr(97+degree)}' value"))
            f_array.append(constant)
        elif i == 1:
            linear = int(input(f"input '{chr(97+degree-1)}x' value"))
            f_array.append(str(linear)+"x+")
        else:
            polynomial = int(input(f"input '{chr(97+degree-i)}x^{i}' value"))
            f_array.append(str(polynomial)+f"x^{i}+")
        print(f_array)
    function = ''.join(map(str, f_array))
    print(function)
f()
