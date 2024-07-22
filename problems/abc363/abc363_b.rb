
def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, t, p = getis

l = getis
l.sort!

if l[-p] > t
  puts 0
  exit
end

puts t - l[-p]