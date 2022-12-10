import re

with open("7_input.txt") as f:
    lines = f.readlines()
    dirs = []
    dir_sizes = {}
    for l in lines:
        if l.startswith("$"):
            command = l.strip().split(" ")
            if command[1] == "cd":
                if command[2] == "..":
                    dirs.pop()
                else:
                    dirs.append(command[2])
                continue

        size = re.findall('\d+', l)
        if len(size) > 0:
            folder = "/"
            for d in dirs:
                if d != "/":
                    folder += d + "/"
                else:
                    folder = "/"
                dir_sizes[folder] = dir_sizes.get(folder, 0) + int(size[0])

    part1 = 0
    for v in dir_sizes.values():
        if v <= 100000:
            part1 += v
    print(f"Part 1: {part1}")

    needed_space = dir_sizes["/"] - 40000000
    part2 = sorted([v for v in dir_sizes.values() if v >= needed_space])[0]
    print(f"Part 2: {part2}")
