# Index sums
win_sums = [2, 1, 3]

index_lookup = {
    "A": 0,
    "B": 1,
    "C": 2,
    "X": 0,
    "Y": 1,
    "Z": 2,
}

# Part 1
with open("2_input.txt") as f:
    score = 0
    for line in f.readlines():
        res = line.split(" ")
        opp = index_lookup[res[0].rstrip()]
        me = index_lookup[res[1].rstrip()]

        if me + opp == win_sums[me]:
            score += 6
        elif me == opp:
            score += 3

        score += me + 1

    print(f"Part 1: {score}")

# Part 2
with open("2_input.txt") as f:
    score = 0
    for line in f.readlines():
        res = line.split(" ")
        opp = index_lookup[res[0].rstrip()]
        desired_result = res[1].rstrip()
        if desired_result == "X":
            me = win_sums[opp] - opp
            score += me + 1
        if desired_result == "Y":
            score += opp + 1 + 3
        if desired_result == "Z":
            for i in range(0, 3):
                if i + opp == win_sums[i]:
                    score += i + 1 + 6

    print(f"Part 2: {score}")
