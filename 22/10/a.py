
# with open("input.txt", "r") as archivo:
with open("test.txt", "r") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]

breaks1 = [20, 60, 100, 140, 180, 220]
breaks2 = [40, 80, 120, 160, 200, 240]

rows = [ [] for _ in range(6) ]
cycles = 0
total  = 0
X = 1 

while len(lineas) != 0:
    curr = lineas.pop(0).split()

    for i in range(len(curr)):
        cycles += 1

        print( '#' if (cycles - (cycles//40) * 40) in range(X, X+3) else '.', end="")
        # print( '#' if cycles in range(X, X+3) else '.', end="")
        if cycles > 180:
            print(" ", cycles, (cycles - (cycles//40) * 40), X)

        if cycles in breaks2:
            print()
            total += X * cycles

    if len(curr) == 2:
        X += int(curr[-1]) 

print(total)
