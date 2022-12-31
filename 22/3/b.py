
# 2789 TO HIGH

score = lambda x: ord(x.lower()) - ord('a') + x.isupper()*26 + 1

with open("input.txt") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]

N = 3
total = 0
for i in range(0,len(lineas),N):
    #a = set(lineas[i:i+N][:len(lineas[i:i+N])//2])
    #a = [ set(j[:len(j)//2]) for j in lineas[i:i+N] ]
    #b = set(lineas[i:i+N][len(lineas[i:i+N])//2:])
    #b = [ set(j[len(j)//2:]) for j in lineas[i:i+N] ]
    c = [ sorted(set(lineas[j])) for j in range(i,i+N) ]

    for j in c[0]:
        if sum([ j in k for k in c[1:] ]) == N-1:
            total += score(j)

print(total)
