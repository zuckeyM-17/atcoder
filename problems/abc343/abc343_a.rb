def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

a, b = getis

sum = a + b

puts 9 - a - b