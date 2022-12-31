
with open("input.txt", "r") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]


#lineas = '''30373
#25512
#65332
#33549
#35390'''.split("\n")

lineas   = [ list(map(int,i))          for i in lineas ]
visibles = [ [ False for _ in lineas ] for _ in lineas ]
scenic   = [ [ 1     for _ in lineas ] for _ in lineas ]
ideal_score = -1

N = len(lineas)

for i in range(N):
    range_r = [ 0 for i in range(10) ]
    range_l = [ 0 for i in range(10) ]
    range_u = [ 0 for i in range(10) ]
    range_d = [ 0 for i in range(10) ]

    for j in range(N):
        scenic[i][j]     *= range_r[lineas[i][j]] + \
                int( range_r[lineas[i][j]] != j)

        scenic[i][N-j-1] *= range_l[lineas[i][N-j-1]] + \
                int( range_l[lineas[i][N-j-1]] != j)

        scenic[j][i]     *= range_u[lineas[j][i]] + \
                int( range_u[lineas[j][i]] != j)

        scenic[N-j-1][i] *= range_d[lineas[N-j-1][i]] + \
                int( range_d[lineas[N-j-1][i]] != j)

        visibles[i][j]     |= range_r[lineas[i][j]]     == j
        visibles[i][N-j-1] |= range_l[lineas[i][N-j-1]] == j

        visibles[j][i]     |= range_u[lineas[j][i]] == j
        visibles[N-j-1][i] |= range_d[lineas[N-j-1][i]] == j

        for k in range(10):
            range_r[k] = range_r[k] + 1 if lineas[i][j]     < k else 0
            range_l[k] = range_l[k] + 1 if lineas[i][N-j-1] < k else 0

            range_u[k] = range_u[k] + 1 if lineas[j][i]     < k else 0
            range_d[k] = range_d[k] + 1 if lineas[N-j-1][i] < k else 0

print(max(sum(scenic, [])))
