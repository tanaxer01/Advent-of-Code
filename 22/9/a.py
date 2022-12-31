
with open("input.txt", "r") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]

    hx, hy = (0,0)
    tx, ty = (0,0)

    for i in lineas:
        #print(i, "|", hx,hy, "|", tx, ty)
        if i[0]   == 'U':
            hx += int(i[2])

        elif i[0] == 'D':
            hx -= int(i[2])

        elif i[0] == 'R':
            hy += int(i[2])

        elif i[0] == 'L':
            hy -= int(i[2])

        grid = [[ '.' for _ in range(min(15, max(5, hx))) ] for _ in range(min(15, max(5, hy)))]
        grid[0][0] = "T"
        grid[0][0] = "H"

        print("".join([j for j in grid[0]]), end="\r")

