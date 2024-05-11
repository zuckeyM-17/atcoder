def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

a = getis
b = getis

puts a.sum - b.sum + 1
