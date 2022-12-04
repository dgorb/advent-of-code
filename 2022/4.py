with open("4_input.txt") as f:
    assignments = f.readlines()
    pairs = 0
    overlap = 0
    for a in assignments:
        elfA, elfB = a.split(",")
        rangeA, rangeB = [[int(a), int(b)] for (a, b) in [elfA.split("-"), elfB.split("-")]]
        setA = set([n for n in range(rangeA[0], rangeA[1]+1)])
        setB = set([n for n in range(rangeB[0], rangeB[1]+1)])
        if setA.issubset(setB) or setB.issubset(setA):
            pairs += 1
        if len(setA.intersection(setB)) > 0 or len(setB.intersection(setA)) > 0:
            overlap += 1

    print(f"Part 1: {pairs}")
    print(f"Part 2: {overlap}")
