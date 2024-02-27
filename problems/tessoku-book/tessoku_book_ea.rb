def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

m = Hash.new(0)
n.times do
  a = geti
  m[a] += 1
end

puts m.values.reduce(0) { |c, v| c + v * (v-1)/2 }