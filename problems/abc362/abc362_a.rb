def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

r, g, b = getis
c = gets.chomp


def min(a, b)
  a < b ? a : b
end

if c == 'Red'
  puts min(g, b)
elsif c == 'Green'
  puts min(r, b)
else
  puts min(r, g)
end