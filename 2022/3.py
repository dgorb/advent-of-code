def item_priority(item):
    if item.isupper():
        return ord(item) - 38
    else:
        return ord(item) - 96

with open("3_input.txt") as f:
    backpacks = f.readlines()

    s = 0
    for contents in backpacks:
        contents = contents.rstrip()
        comp1 = set(contents[:len(contents)//2])
        comp2 = set(contents[len(contents)//2:])
        shared = comp1.intersection(comp2)
        for i in shared:
            s += item_priority(i)

    print(f"Part 1: {s}")

    s = 0
    for i in range(0, len(backpacks), 3):
        elf1, elf2, elf3 = [b.rstrip() for b in backpacks[i:i+3]]
        badge = set(elf1).intersection(set(elf2)).intersection(set(elf3))
        s += item_priority(badge.pop())

    print(f"Part 2: {s}")
