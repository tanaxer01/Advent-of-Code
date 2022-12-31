
with open("input.txt", "r") as archivo:
    lineas = [ i.strip() for i in archivo.readlines() ]


next_line = lambda: lineas.pop(0).lstrip("$ ") if len(lineas) != 0 else None

pwd, tree, weights = "", {}, {}
while (curr := next_line()) != None: 
    if   curr[:2] == "cd":
        print(curr, pwd, sep="-")
        if pwd == "":
            pwd = curr[3:]
            continue
        elif curr[3:] != "..":
            tree[curr[3:]] = pwd
            pwd = curr[3:]
        else:
            print(pwd, tree[pwd])
            pwd = tree[pwd]
            break

    elif curr[:2] == "ls":
        pass
    else:
        pass
       

