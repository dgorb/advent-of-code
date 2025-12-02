local function solve()
    local current = 50
    local part_1 = 0
    local part_2 = 0

    for line in io.lines("01.txt") do
        local direction, amount = line:match("(%a)(%d+)")
        amount = tonumber(amount)

        if direction == "R" then
            part_2 = part_2 + math.floor((current + amount) / 100)
            current = (current + amount) % 100
        elseif direction == "L" then
            if current == 0 then
                -- Edge case: starting at 0, only count complete rotations
                part_2 = part_2 + math.floor(amount / 100)
            elseif amount >= current then
                part_2 = part_2 + 1 + math.floor((amount - current) / 100)
            end
            current = (current - amount) % 100
        end

        if current == 0 then
            part_1 = part_1 + 1
        end
    end

    print(part_1)
    print(part_2)
end

solve()
