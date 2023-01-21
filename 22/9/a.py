
N = 10

with open("input.txt", "r") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]
    #lineas = "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2".split("\n")
    #lineas = "R 5\nU 8\nL 8\nD 3\nR 17\nD 10\nL 25\nU 20".split("\n")

movs  = { 'U': (0, +1), 'D': (0, -1), 'R': (+1, 0), 'L': (-1, 0) }
nudos = [ [0,0] for i in range(N) ]

visited = set()
visited.add(tuple(nudos[-1]))

for num, line in enumerate(lineas):
    direction, steps = line.split()
    for _ in range( int(steps) ):
        nudos[0][0] += movs[direction][0]
        nudos[0][1] += movs[direction][1]

        for i in range(N-1):
            dx = nudos[i][0] - nudos[i+1][0]
            dy = nudos[i][1] - nudos[i+1][1]

            if abs(dx) > 1 or abs(dy) > 1:
                nudos[i+1][0] += 1 if dx > 0 else -1 if dx < 0 else 0
                nudos[i+1][1] += 1 if dy > 0 else -1 if dy < 0 else 0

            visited.add(tuple(nudos[-1]))

print( len(visited) )
