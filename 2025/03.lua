local function solve()
    local part_1 = 0
    local part_2 = 0

    for line in io.lines("03.txt") do
        local joltages = {}

        for i = 1, #line do
            table.insert(joltages, tonumber(string.sub(line, i, i)))
        end

        local largest = 0
        for i = 1, #joltages do
            for j = i + 1, #joltages do
                local sum = joltages[i] * 10 + joltages[j]
                if sum and sum > largest then
                    largest = sum
                end
            end
        end

        part_1 = part_1 + largest
    end

    print("Part 1:", part_1)
end

solve()
