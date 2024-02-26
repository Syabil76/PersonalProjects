#to be paired with my functions code

def integration():
    bound = int(input("please input your boundaries"))
    integrated = []
    newpoly = []
    for i in range(degree+1, 0, -1):
        print(f"{funny[degree+1-i]} to the power of {i} divided by {i}")
        new = round(
            ((funny[degree+1-i])/i), 5
        )
        integrated.append(new)
        newpoly.append(''.join(map(str, new))+"+")
    print(f"indefinite integral = {''.join(map(str, newpoly))} + C")


integration()
