def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti
s = gets.chomp
q = geti


h = "abcdefghijklmnopqrstuvwxyz".chars.map do |c|
	[c, c]
end.to_h

q.times do
	c, d = gets.chomp.split(" ")
	h.each do |k, v|
		if v == c
			h[k] = d
		end
	end
end

s.chars.each do |c|
	print h[c]
end

puts ""