local removeToken = function(type, name)
    local ns = "token:" .. type .. ":"
    local token = redis.call("GET", ns .. name)
    if token ~= false then
        redis.call("DEL", ns .. token)
        redis.call("DEL", ns .. name)
    end
    return true
end

return removeToken(KEYS[1], KEYS[2])