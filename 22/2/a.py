
WIN_HAND = ['C X', 'A Y', 'B Z']
DRW_HAND = ['A X', 'B Y', 'C Z']
LST_HAND = ['B X', 'C Y', 'A Z']

WIN_HAND2 = ['C Z', 'A Z', 'B Z']
DRW_HAND2 = ['A Y', 'B Y', 'C Y']
LST_HAND2 = ['B X', 'C X', 'A X']

with open("input.txt", "r") as archivo:
    lineas  = archivo.read()[:-1]

    jugadas = lineas.split("\n")

    score = score2 = 0
    for i in jugadas:
        # Item 1
        if i in WIN_HAND:
            score  += 6 + WIN_HAND.index(i) + 1
        elif i in DRW_HAND:
            score += 3 + DRW_HAND.index(i) + 1
        else:
            score += LST_HAND.index(i) + 1

        # Item 2
        if i in WIN_HAND2:
            score2  += 6 + WIN_HAND2.index(i) + 1
        elif i in DRW_HAND2:
            score2 += 3 + DRW_HAND2.index(i) + 1
        else:
            score2 += LST_HAND2.index(i) + 1

    print(f"1. {score}")
    print(f"2. {score2}")


