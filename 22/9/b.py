import numpy as np

def part_one():
    with open("input.txt", "r") as archivo:
        lineas = [ i.strip() for i in archivo.readlines() ]

    head = (0, 0)
    tail = (0, 0)
    visited = set()
    visited.add(tail)
    print(visited)

    for motion in lineas:
        direction, steps = motion.split()
        for _ in range(int(steps)):
            if direction == "U":
                head = (head[0], head[1] + 1)
            elif direction == "D":
                head = (head[0], head[1] - 1)
            elif direction == "R":
                head = (head[0] + 1, head[1])
            elif direction == "L":
                head = (head[0] - 1, head[1])
            diff_x = head[0] - tail[0]
            diff_y = head[1] - tail[1]
            if abs(diff_x) > 1 or abs(diff_y) > 1:
                tail = (tail[0] + np.sign(diff_x), tail[1] + np.sign(diff_y))
                visited.add(tail)
    return len(visited)

def solution(length):
    with open("input.txt", "r") as archivo:
        lineas = [ i.strip() for i in archivo.readlines() ]
    lineas = "R 5\nU 8\nL 8\nD 3\nR 17\nD 10\nL 25\nU 20".split("\n")

    movs  = { 'U': (0, +1), 'D': (0, -1), 'R': (+1, 0), 'L': (-1, 0) }
    knots = [ (0,0) for i in range(length) ]

    visited = set()
    visited.add(knots[-1])

    for num, motion in enumerate(lineas):
        direction, steps = motion.split()
        for _ in range(int(steps)):
            knots[0] = (knots[0][0] + movs[direction][0],
                        knots[0][1] + movs[direction][1])
            for i in range(length-1):
                dx = knots[i][0] - knots[i+1][0]
                dy = knots[i][1] - knots[i+1][1]

                if abs(dx) > 1 or abs(dy) > 1:
                    knots[i+1] = (knots[i+1][0] + np.sign(dx),
                                  knots[i+1][1] + np.sign(dy))

                visited.add(tuple(knots[-1]))

        print(visited)


    return len(visited)

print(solution(10))
