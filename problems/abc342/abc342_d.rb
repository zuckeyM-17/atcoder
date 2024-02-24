def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti
a = getis

ans = 0
h = {}
zeros = 0
a.each.with_index do |i, idx|
	if i == 0
		ans += n-1 - zeros
		zeros+=1
		next
	end
	(2..).each do |x|
		x2 = x * x
		i /= x2 while i % x2 == 0
		if i < x2
			h[i] == nil ? h[i] = 1 : h[i] += 1
			break
		end
	end
end

h.values.each do |x|
	if x < 2
		next
	end
	ans += (x * (x - 1)) / 2
end

puts ans