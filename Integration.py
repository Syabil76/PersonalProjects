#to be paired with my functions code

def integration():
  # bounds = int(input("please input your boundaries"))
  integrated = []
  for i in range(degree+1, 0, -1):
      print(f"{funny[degree+1-i]} to the power of {i} divided by {i}")
      new = (
          (((funny[degree+1-i])**i)/i)
      )
      integrated.append(new)
  print(integrated)

integration()
