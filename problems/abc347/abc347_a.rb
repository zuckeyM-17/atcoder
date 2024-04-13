def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, k = getis
a = getis

puts a.select { |x| x % k == 0 }.map { |x| x / k }.join(" ")