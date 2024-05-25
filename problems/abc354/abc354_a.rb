def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

h = geti

i = 0
cur = 0
loop do
  i += 1
  cur += 2 ** i
  break if cur >= h
end

puts i + 1