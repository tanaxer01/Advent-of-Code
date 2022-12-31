

with open("input.txt", "r") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]

#lineas = '''30373
#25512
#65332
#33549
#35390'''.split("\n")

lineas = [ list(map(int,i)) for i in lineas ]
is_visible = [ [ False for i in lineas[1:-1] ] for j in lineas[1:-1] ]

# HORIZONTAL CHECK
for i in range( 1, len(lineas)-1 ):
    # LEFT - RIGHT
    tallest = lineas[i][0]
    for j in range( 1, len(lineas)-1 ):
        is_visible[i-1][j-1] |= lineas[i][j] > tallest

        tallest = lineas[i][j] if lineas[i][j] > tallest else tallest

    # RIGHT - LEFT
    tallest = lineas[i][-1]
    for j in range( len(lineas)-2, 0, -1 ):
        is_visible[i-1][j-1] |= lineas[i][j] > tallest

        tallest = lineas[i][j] if lineas[i][j] > tallest else tallest


# VERTICAL CHECK
for j in range( 1, len(lineas)-1 ):
    # UP - DOWN
    tallest = lineas[0][j]
    for i in range( 1, len(lineas)-1 ):
        is_visible[i-1][j-1] |= lineas[i][j] > tallest

        tallest = lineas[i][j] if lineas[i][j] > tallest else tallest

    # DOWN - UP
    tallest = lineas[-1][j]
    for i in range( len(lineas)-2, 0, -1 ):
        is_visible[i-1][j-1] |= lineas[i][j] > tallest

        tallest = lineas[i][j] if lineas[i][j] > tallest else tallest


print( len(lineas)*4 - 4+ sum(sum(is_visible, [])) )
