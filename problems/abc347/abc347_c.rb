def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, a, b = getis
d = getis

q = d.map do |x|
  x % (a + b)
end

puts q.max - q.min < a ?  "Yes" : "No"
