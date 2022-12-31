
with open("input.txt", "r") as archivo:
    lineas = archivo.read()[:-1]

    elfos  =  [ sum(map(int,i.split("\n"))) for i in lineas.split("\n\n") ]

    sorted_elfos = sorted(elfos, reverse=True)
    print(f"1. {sorted_elfos[0]}")
    print(f"2. {sum(sorted_elfos[:3])}")


