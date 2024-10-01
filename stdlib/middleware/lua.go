package middleware

const (
	ResetScript = `	
		local routeKey = KEYS[1]
		local staticKey = KEYS[2]
		local routeDeadline = ARGV[1]

		redis.call('HSET', staticKey, "Count", 1)
		redis.call('HSET', routeKey, "Count", 1, "Deadline", routeDeadline)

		return 0
	`
	Script = `
		local result = {}
		local routeKey = KEYS[1]
		local staticKey = KEYS[2]

		local routeLimit = tonumber(ARGV[1])
		local staticLimit = tonumber(ARGV[2])
		local routeDeadline = tonumber(ARGV[3])
		local now = tonumber(ARGV[4])

		local staticCount = redis.call('HGET', staticKey, "Count")

		-- First time visit
		if not staticCount then 
			redis.call('HSET', staticKey, "Count", 1)
			redis.call('HSET', routeKey, "Count", 1, "Deadline", routeDeadline)
			result[1] = staticLimit - 1
			result[2] = routeLimit - 1
			result[3] = routeDeadline 

			return result
		end 

		local routeInfo = redis.call('HGETALL', routeKey)

		if #routeInfo == 0 then 
			if tonumber(staticCount) < staticLimit then
				result[1] = staticLimit - redis.call('HINCRBY', staticKey, "Count", 1)
			else 
				result[1] = -1
			end

			redis.call('HSET', routeKey, "Count", 1, "Deadline", routeDeadline)
			result[2] = routeLimit - 1
			result[3] = routeDeadline

			return result
		end

		local rDead = tonumber(routeInfo[4]) -- expired time	
		local rCount = tonumber(routeInfo[2])  
		local sCount = tonumber(staticCount)
		
		if rDead < now then
			if tonumber(staticCount) < staticLimit then 
				result[1] = staticLimit - redis.call('HINCRBY', staticKey, "Count", 1)
			else 
				result[1] = -1
			end
			
			redis.call('HSET', routeKey, "Count", 1, "Deadline", routeDeadline)
			result[2] = routeLimit - 1
			result[3] = routeDeadline

			return result
		end

		if sCount < staticLimit then 
			result[1] = staticLimit - redis.call('HINCRBY', staticKey, "Count", 1)
		else 
			result[1] = -1
		end

		if rCount < routeLimit then 
			result[2] = redis.call('HINCRBY', routeKey, "Count", 1)
		else 
			result[2] = -1
		end

		result[3] = rDead
		
		return result
	`
)
