local function solve()
    local current = 50
    local part_1 = 0

    for line in io.lines("01.txt") do
        local direction, amount = line:match("(%a)(%d+)")

        if direction == "R" then
            current = (current + amount) % 100
        elseif direction == "L" then
            current = (current - amount) % 100
        end

        if current == 0 then
            part_1 = part_1 + 1
        end
    end

    print(part_1)
end

solve()
