
with open("input.txt", "r") as archivo:
    lineas   = archivo.readlines()

    starting = [ i.strip() for i in lineas[:8] ]
    lineas   = lineas[10:]

    queues  = [ [j[i] for j in starting[::-1] if j[i] != ' ' ] for i in range(1, len(starting[0]), 4) ]
    queues2 = [ [j[i] for j in starting[::-1] if j[i] != ' ' ] for i in range(1, len(starting[0]), 4) ]
    
    for i in lineas:
        cant, tl   = i[5:].strip().split(" from ")
        a, b = tl.split(" to ")

        queues[int(b)-1] += queues[int(a)-1][-1*int(cant):][::-1]
        queues[int(a)-1] = queues[int(a)-1][:-1*int(cant)]

        queues2[int(b)-1] += queues2[int(a)-1][-1*int(cant):]
        queues2[int(a)-1] = queues2[int(a)-1][:-1*int(cant)]


    print(f"1. {''.join([ i[-1] for i in queues  ])}")
    print(f"2. {''.join([ i[-1] for i in queues2 ])}")

