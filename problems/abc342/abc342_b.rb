def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

p = getis

q = geti

q.times do
	a, b = getis

	if p.index(a) < p.index(b)
		puts a
	else
		puts b
	end
end
