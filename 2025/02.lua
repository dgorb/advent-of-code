local function split_by_comma(str)
    local result = {}
    for match in str:gmatch("([^,]+)") do
        table.insert(result, match)
    end
    return result
end

local function is_repeated_twice(str)
    local len = #str

    if len % 2 ~= 0 then
        return false, nil
    end

    local pattern = string.sub(str, 1, len // 2)
    local doubled = string.rep(pattern, 2)

    return str == doubled, pattern
end

local function has_repeating_pattern(str)
    local len = #str

    for pattern_len = 1, len // 2 do
        if len % pattern_len == 0 then
            local pattern = string.sub(str, 1, pattern_len)
            local repeated = string.rep(pattern, len // pattern_len)

            if repeated == str then
                return true
            end
        end
    end

    return false
end

local function solve()
    local part_1 = 0
    local part_2 = 0

    local line = io.lines("02.txt")()
    for _, id_range in ipairs(split_by_comma(line)) do
        local first_id, last_id = id_range:match("([^-]+)-([^-]+)")

        for id = math.tointeger(first_id), math.tointeger(last_id) do
            if is_repeated_twice(tostring(id)) then
                part_1 = part_1 + id
            end

            if has_repeating_pattern(tostring(id)) then
                part_2 = part_2 + id
            end
        end
    end

    print("Part 1:", part_1)
    print("Part 2:", part_2)
end

solve()
