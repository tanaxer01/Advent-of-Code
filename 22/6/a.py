
with open("input.txt", "r") as archivo:
    start = archivo.read().strip()

    item1 = start
    len1 = len(item1) + 4
    while len(set(item1[:7])) != 7 and len(item1) > 0:
        item1 = item1[1:]

    print(f"1. {len1 - len(item1)}")
    # Item 2
    linea2 = start
    starting_len2 = len(linea2) + 14
    while len(set(linea2[:14])) != 14 and len(linea2) > 0:
        linea2 = linea2[1:]


    print(f"2. {starting_len2 - len(linea2)}")


