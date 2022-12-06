def marker_pos(signal, marker_length):
    chars = []
    for i, char in enumerate(signal):
        if len(chars) == marker_length:
            chars.pop(0)
        chars.append(char)

        if len(set(chars)) == marker_length:
            return i + 1

with open("6_input.txt") as f:
    signal = f.readline().rstrip()

    print(f"Part 1: {marker_pos(signal, 4)}")
    print(f"Part 2: {marker_pos(signal, 14)}")
