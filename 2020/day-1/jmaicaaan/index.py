from data import getContents

list = getContents()
def main():
  product = 0
  for x in list:
    for y in list:
      if x == y:
        continue
      if int(x) + int(y) == 2020:
        product = int(x) * int(y)
        return product
      
print(main())

# day 2

def main2():
  product = 0
  for x in list:
    for y in list:
      for z in list:
        if x == y or x == z:
          continue
        if int(x) + int(y) + int(z) == 2020:
          product = int(x) * int(y) * int(z)
          return product
      
print(main2())