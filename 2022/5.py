import re

stacks = [
    "BQC",
    "RQWZ",
    "BMRLV",
    "CZHVTW",
    "DZHBNVG",
    "HNPCJFVQ",
    "DGTRWZS",
    "CGMNBWZP",
    "NJBMWQFP"
]
stacksB = stacks.copy()

with open("5_input.txt") as f:
    moves = f.readlines()

    for i in range(10, len(moves)):
        m = re.findall('\d+', moves[i])

        stacks[int(m[2])-1] += stacks[int(m[1])-1][-int(m[0]):][::-1]
        stacks[int(m[1])-1] = stacks[int(m[1])-1][:-int(m[0])]

        stacksB[int(m[2])-1] += stacksB[int(m[1])-1][-int(m[0]):]
        stacksB[int(m[1])-1] = stacksB[int(m[1])-1][:-int(m[0])]

    answer1 = ''.join([a[-1] for a in stacks])
    answer2 = ''.join([a[-1] for a in stacksB])
    print(f"Part 1: {answer1}")
    print(f"Part 2: {answer2}")
