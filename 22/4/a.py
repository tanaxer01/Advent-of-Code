
with open("input.txt") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]

total = 0
for i in lineas:
    a, b = [ [ int(k) for k in j.split("-") ] for j in i.split(",") ]

    if a[0] <= b[0] and a[1] >= b[1]: 
        total += 1
    elif a[0] >= b[0] and a[1] <= b[1]: 
        total += 1

print(total)
