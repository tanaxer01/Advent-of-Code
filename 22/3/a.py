
score = lambda x: ord(x.lower()) - ord('a') + x.isupper()*26 + 1

with open("input.txt") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]

total = 0
for i in lineas:
    a = set(i[:len(i)//2])
    b = set(i[len(i)//2:])

    for j in a:
        if j in b:
            total += score(j)

print(total)
