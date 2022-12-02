with open("aoc1_input.txt") as f:
    elves = []
    i = 0

    for line in f.readlines():
        if line == "\n":
            i += 1
            continue

        if len(elves) == i:
            elves.append(0)

        elves[i] += int(line)

    elves = sorted(elves)

    print(f"Part 1: {elves[-1]}")
    print(f"Part 2: {sum(elves[-3:])}")
