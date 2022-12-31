# 937 TO HIGH

with open("input.txt") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]

total = 0
for i in lineas:
    a, b = [ [ int(k) for k in j.split("-") ] for j in i.split(",") ]

    if   a[0] >= b[0] and a[0] <= b[1]:
        total += 1
    elif a[1] >= b[0] and a[1] <= b[1]:
        total += 1
    elif b[0] >= a[0] and b[0] <= a[1]:
        total += 1
    elif b[0] >= a[0] and b[0] <= a[1]:
        total += 1

    break


print(total)
